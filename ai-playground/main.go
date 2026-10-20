package main

import (
	"fmt"
	"log"

	"ai-playground/config"
	"ai-playground/controllers"
	"ai-playground/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Reads OPENAI_API_KEY, AI_BASE_URL, AI_MODEL from .env
	config.LoadDotEnv(".env")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.LoggerMiddleware())

	router.GET("/health", controllers.GetHealth)
	router.POST("/api/chat", controllers.PostChat)

	fmt.Println("server running on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
