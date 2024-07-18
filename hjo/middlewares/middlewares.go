package middlewares


import (
	"net/http"
	"strings"
	"hjo/config"
	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqSecret := c.Request.Header.Get("x-test-key")
		var secret string

		if secret = config.CFG.V.GetString("GO_EASY_KEY"); len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if reqSecret != secret {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}


