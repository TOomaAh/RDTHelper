package middleware

import (
	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/model"
	"github.com/gin-gonic/gin"
)

// Check if cookie is save
func CheckAuthenticated(db *database.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") == "" && c.Query("token") == "" {
			// c response 401 unauthorized not redirect to login
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// extract token from header
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			token = c.Query("token")
		} else {
			//remove Bearer from token
			token = token[7:]
		}

		userId, err := model.GetUserIdFromToken(token)

		if err != nil {
			c.JSON(401, gin.H{"error": "User not found"})
			return
		}

		user, err := db.FindUserByUserId(userId)

		if err != nil {
			c.JSON(401, gin.H{"error": "User not found"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
