package rdthelper

import (
	"github.com/TOomaAh/RDTHelper/internal/api"
	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/internal/middleware"
	"github.com/TOomaAh/RDTHelper/internal/web"
	"github.com/TOomaAh/RDTHelper/pkg/realdebrid"
	"github.com/gin-gonic/gin"
)

func Run(db *database.Database, r *gin.Engine) {

	r.Use(middleware.DatabaseMiddleware(db))

	r.Use(func(c *gin.Context) {
		// get token
		if db.HasUser() {
			token, err := db.GetToken()
			if err != nil {
				return
			}
			realDebridClient := realdebrid.Authentication(token)
			c.Set("rd", realDebridClient)

		}
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	group := r.Group("/api/v1")
	front := r.Group("/web")
	login := r.Group("/")
	login.Use(middleware.CheckUserExist)
	front.Use(middleware.CheckUserExist)

	u := api.RegisterUser(group, db)
	api.NewTorrentApi(group)

	web.NewWebGroup(front)
	web.RegisterLogin(login, u)

	r.GET("/", func(c *gin.Context) {
		//redirect to /web
		c.Redirect(302, "/web/home")
	})

	r.Run()
}
