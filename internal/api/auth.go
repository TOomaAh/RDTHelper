package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Authentication struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterAuth(c *gin.Context) {
	var auth Authentication

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
