package middleware

import (
	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/gin-gonic/gin"
)

// Insert database in context
// TODO: Replace db parameter with Repository pattern
func DatabaseMiddleware(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
