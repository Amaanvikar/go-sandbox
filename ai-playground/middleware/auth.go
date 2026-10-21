package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		apikey := c.GetHeader("X-API-kEY") // Get the header value

		expectedKey := "abc123" // dummy hardcoded key value

		// check if the string is empty
		if apikey == "" {

			// abort the request and return 401 unauthorized error
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "401 unauthorized: Missing API Key",
			})

			return // stops execution of this middleware
		}

		if apikey != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "401 unauthorized: Wrong API Key",
			})
			return
		}

		c.Next() // process with the next handler
	}
}
