package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // records the starting time

		c.Next() // continue processing the request

		elapsed := time.Since(start) // calcualte the elapsed time

		// extract the response statuscode
		statusCode := c.Writer.Status()

		// Extract the request method and path
		method := c.Request.Method
		path := c.Request.URL.Path

		fmt.Printf("API logs | Method: %s | Path: %s | Status: %d | Duration: %v\n",
			method, path, statusCode, elapsed)
	}
}
