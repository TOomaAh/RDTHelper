package middleware

import (
	"github.com/gin-gonic/gin"
)

// Check if cookie is save
func CheckAuthenticated(c *gin.Context) {
	// if token is not present in header redirect to login
	if c.Request.Header.Get("Authorization") == "" && c.Request.URL.Path != "/login" {
		c.Redirect(302, "/login")
		return
	}

	c.Next()
}
