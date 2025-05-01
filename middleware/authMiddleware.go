package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userRole := session.Get("role")
		if userRole != role {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		if role != "" && role != userRole {
			c.String(http.StatusForbidden, "Access denied !")
			c.Abort()
			return
		}
		c.Next()
	}
}
