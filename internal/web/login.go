package web

import (
	"github.com/TOomaAh/RDTHelper/internal/api"
	"github.com/gin-gonic/gin"
)

type LoginError struct {
	Logout bool
	Err    bool
}

func RegisterLogin(group *gin.RouterGroup, u *api.UserWebGroup) {

	var loginError LoginError = LoginError{Logout: false, Err: false}

	group.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{
			"logout": loginError.Logout,
			"err":    loginError.Err,
		})
	})

	group.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	})

	group.POST("/perform_signup", func(c *gin.Context) {
		u.PerformSignup(c)
	})
}
