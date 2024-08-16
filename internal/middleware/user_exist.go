package middleware

import (
	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/gin-gonic/gin"
)

// Check if one user exist middleware
func CheckUserExist(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db.HasUser() && c.Request.URL.Path == "/signup" {
			c.Redirect(302, "/login")
			return
		}
		if !db.HasUser() && c.Request.URL.Path != "/signup" {
			c.Redirect(302, "/signup")
			return
		}
		c.Next()
	}
}
