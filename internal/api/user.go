package api

import (
	"fmt"

	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/internal/middleware"
	"github.com/TOomaAh/RDTHelper/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserWebGroup struct {
	db *database.Database
}

func RegisterUser(group *gin.RouterGroup, db *database.Database) *UserWebGroup {
	u := UserWebGroup{db}
	group.POST("/signup", u.signup)
	group.POST("/login", u.login)

	userGroup := group.Group("/")
	userGroup.Use(middleware.CheckAuthenticated(db))
	userGroup.GET("/settings", u.settings)

	return &u
}

func (u *UserWebGroup) settings(c *gin.Context) {
	if user, exist := c.Get("user"); exist {
		c.JSON(200, user.(*model.User).ToDTO())
		return
	}
	c.JSON(400, gin.H{
		"error": "User not found",
	})

}

func (u *UserWebGroup) signup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Create(db)
	c.JSON(200, gin.H{
		"user": user,
	})
}

func (u *UserWebGroup) verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserWebGroup) login(c *gin.Context) {
	var input Authentication

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("input: %v\n", input)

	if input.Username == "" {
		c.JSON(400, gin.H{"error": "Username is required"})
		return
	}

	user, err := u.db.FindUserByUsername(input.Username)

	if err != nil {
		c.JSON(400, gin.H{"error": "Username or password incorrect"})
		return
	}

	err = u.verifyPassword(input.Password, user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "Username or password incorrect"})
		return
	}

	token, err := model.GenerateToken(user.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Error while generating token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}
