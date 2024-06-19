package middleware

import (
	"fmt"

	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/gin-gonic/gin"
)

// Check if one user exist middleware
func CheckUserExist(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	fmt.Println(db.HasUser())
	if !db.HasUser() && c.Request.URL.Path != "/signup" {
		c.Redirect(302, "/signup")
		return
	}
	c.Next()
}
