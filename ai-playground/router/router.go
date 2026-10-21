package router

import (
	"ai-playground/controllers"
	"ai-playground/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.LoggerMiddleware())

	router.GET("/health", controllers.GetHealth)
	// router.POST("/api/chat", controllers.PostChat)
	router.POST("/api/chat", middleware.APIKeyMiddleware(), controllers.PostChat)

	return router

}
