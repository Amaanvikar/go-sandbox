package model

// ChatRequest is what the browser / client sends to our API.
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

// ChatResponse is what our API sends back.
type ChatResponse struct {
	Reply string `json:"reply"`
}
