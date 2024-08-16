package web

import (
	"github.com/TOomaAh/RDTHelper/internal/api"
	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterLogin(group *gin.RouterGroup, u *api.UserWebGroup, db *database.Database) {
	group.Use(middleware.CheckUserExist(db))

	group.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})

	group.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	})
}
