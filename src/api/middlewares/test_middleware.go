package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		if apiKey == "1" {
			c.Next()
		}else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "api key is required or invalid",
			})
			return
		}


	}
}
