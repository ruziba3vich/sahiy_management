package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/sahiy_management/internal/application/auth"
)

func RequirePrivilege(authService *auth.Service, resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		hasPrivilege, err := authService.HasPrivilege(c.Request.Context(), userID.(int64), role.(int), resource, action)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check privilege"})
			c.Abort()
			return
		}

		if !hasPrivilege {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient privileges"})
			c.Abort()
			return
		}

		c.Next()
	}
}
