package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // records the starting time

		c.Next() // continue processing the request

		requestID, exists := c.Get("requestID")

		elapsed := time.Since(start) // calcualte the elapsed time

		// extract the response statuscode
		statusCode := c.Writer.Status()

		// Extract the request method and path
		method := c.Request.Method
		path := c.Request.URL.Path

		fmt.Printf("RequestID: | RequestID: %v | Method: %s | Path: %s | StatusCode: %d | Duration: %v | Exist: %t\n",
			requestID, method, path, statusCode, elapsed, exists)
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// generate a unique id
		uniqueID := uuid.New().String()

		// store it in the gin context using c.Header
		c.Set("requestID", uniqueID)

		// adding it to the response header
		c.Header("X-Request-ID", uniqueID)

		// continue requests
		c.Next()
	}
}
