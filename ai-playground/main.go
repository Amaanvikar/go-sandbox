package main

import (
	"fmt"
	"log"

	"ai-playground/config"
	"ai-playground/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Reads OPENAI_API_KEY, AI_BASE_URL, AI_MODEL from .env
	config.LoadDotEnv(".env")

	gin.SetMode(gin.ReleaseMode)

	r := router.SetupRouter()

	fmt.Println("server running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
