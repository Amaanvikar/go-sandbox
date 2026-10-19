package main

import (
	"file-system/controllers"
	"file-system/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	router.GET("/api/health", controllers.GetHealth)

	router.POST("/upload", controllers.UploadFile)
	router.GET("/files", controllers.GetUpload)

	fmt.Println("Server started at :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	router := gin.Default()

// 	router.GET("/health", controllers.GetHealth)

// 	fmt.Println("Server started at: 8080")

// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatal(err)
// 	}
// }
