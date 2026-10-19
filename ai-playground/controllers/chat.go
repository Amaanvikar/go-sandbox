package controllers

import (
	"net/http"

	"ai-playground/model"
	"ai-playground/services"

	"github.com/gin-gonic/gin"
)

// PostChat handles POST /api/chat
// 1) Read JSON from the client
// 2) Ask the AI service for a reply
// 3) Send JSON back
func PostChat(c *gin.Context) {
	var req model.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "send JSON like {\"message\": \"hello\"}"})
		return
	}

	reply, err := services.GetAIReply(req.Message)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.ChatResponse{Reply: reply})
}
