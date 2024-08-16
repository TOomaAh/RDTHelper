package rdthelper

import (
	"github.com/TOomaAh/RDTHelper/internal/api"
	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/internal/web"
	"github.com/gin-gonic/gin"
)

func Run(db *database.Database, r *gin.Engine) {

	r.LoadHTMLGlob("public/templates/*")
	r.Static("/static", "./public/static")

	group := r.Group("/api/v1")
	front := r.Group("/web")
	login := r.Group("/")

	u := api.RegisterUser(group, db)
	api.NewTorrentApi(group, db)

	web.NewWebGroup(front)
	web.RegisterLogin(login, u, db)

	r.GET("/", func(c *gin.Context) {
		//redirect to /web
		c.Redirect(302, "/web/home")
	})

	r.Run()
}
