package middlerwares

import (
	"github.com/danhbuidcn/go_backend_api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
