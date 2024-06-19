package api

import (
	"fmt"

	"github.com/TOomaAh/RDTHelper/internal/database"
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
	group.GET("/users", u.listUsers)
	group.POST("/users", u.createUser)
	group.POST("/signup", u.signup)
	group.POST("/login", u.login)
	return &u
}

func (u *UserWebGroup) listUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	users := model.FindAllUsers(db)
	c.JSON(200, gin.H{
		"users": users,
	})
}

func (u *UserWebGroup) PerformSignup(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	var user model.User
	user.Username = c.PostForm("username")
	//encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	user.RdtToken = c.PostForm("rdt_token")
	if db.CreateUser(&user) != nil {
		c.JSON(400, gin.H{"error": "Error while creating user"})
		return

	}
	c.Redirect(302, "/web/home")
}

func (u *UserWebGroup) PerformLogin(username string, password string) (string, error) {

	user, err := u.db.FindUserByUsername(username)

	if err != nil {
		fmt.Println("user not found")
		return "", fmt.Errorf("user not found")
	}

	err = u.verifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}
	token, err := model.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserWebGroup) createUser(c *gin.Context) {
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

	username := input.Username

	if username == "" {
		c.JSON(400, gin.H{"error": "Username is required"})
		return
	}

	password := input.Password

	user, err := u.db.FindUserByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": "Username incorrect"})
		return
	}

	err = u.verifyPassword(password, user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "Password incorrect"})
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
