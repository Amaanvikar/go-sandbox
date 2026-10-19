package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// openAIMessage matches the shape OpenAI-compatible APIs expect.
type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIRequest struct {
	Model    string          `json:"model"`
	Messages []openAIMessage `json:"messages"`
}

type openAIResponse struct {
	Choices []struct {
		Message openAIMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// GetAIReply talks to an AI provider, or returns a mock reply if no API key is set.
// This keeps the playground usable while you learn — no paid key required at first.
func GetAIReply(userMessage string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Beginner-friendly path: practice the full app flow without an external AI.
	if apiKey == "" {
		return mockReply(userMessage), nil
	}

	return callOpenAI(apiKey, userMessage)
}

func mockReply(userMessage string) string {
	trimmed := strings.TrimSpace(userMessage)
	return fmt.Sprintf(
		"Mock AI (no OPENAI_API_KEY set). You said: %q. "+
			"Set OPENAI_API_KEY in a .env file to talk to a real model.",
		trimmed,
	)
}

func callOpenAI(apiKey, userMessage string) (string, error) {
	// Defaults work with OpenAI. Point AI_BASE_URL at any compatible API if you want.
	baseURL := envOr("AI_BASE_URL", "https://api.openai.com/v1")
	model := envOr("AI_MODEL", "gpt-4o-mini")

	body := openAIRequest{
		Model: model,
		Messages: []openAIMessage{
			{Role: "system", Content: "You are a helpful assistant in a beginner AI playground. Keep answers short and clear."},
			{Role: "user", Content: userMessage},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("could not encode request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+"/chat/completions", bytes.NewReader(jsonBody))
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("AI request failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read AI response: %w", err)
	}

	var parsed openAIResponse
	if err := json.Unmarshal(respBytes, &parsed); err != nil {
		return "", fmt.Errorf("could not parse AI response: %w", err)
	}

	if parsed.Error != nil {
		return "", fmt.Errorf("AI provider error: %s", parsed.Error.Message)
	}
	if len(parsed.Choices) == 0 {
		return "", fmt.Errorf("AI returned no choices (status %d)", resp.StatusCode)
	}

	return parsed.Choices[0].Message.Content, nil
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
