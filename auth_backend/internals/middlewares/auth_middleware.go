package middlewares

import (
	"net/http"
	"strings"

	"github.com/devesh121/userAuth/internals/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from cookie
		token, err := c.Cookie("auth_token")
		if err != nil || strings.TrimSpace(token) == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token not found"})
			c.Abort()
			return
		}

		// Validate token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid token"})
			c.Abort()
			return
		}

		// we can set the user ID / email / role into context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		// Continue to handler
		c.Next()
	}
}
