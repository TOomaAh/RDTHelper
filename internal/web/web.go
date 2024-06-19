package web

import (
	"log"
	"strings"

	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/pkg/realdebrid"
	"github.com/gin-gonic/gin"
)

type WebGroup struct {
	client *realdebrid.RealDebridClient
}

func NewWebGroup(group *gin.RouterGroup) {
	w := &WebGroup{}

	group.Use(func(c *gin.Context) {
		// get rdtClient and catch error
		client, exist := c.Get("rd")
		if !exist {
			return
		}
		w.client = client.(*realdebrid.RealDebridClient)
	})

	group.GET("/settings", func(ctx *gin.Context) {
		db := ctx.MustGet("db").(*database.Database)
		user, _ := db.FindUserByUsername("Thomas")

		ctx.HTML(200, "settings.html", gin.H{
			"username": user.Username,
			"rdtToken": user.RdtToken,
			"password": user.Password,
		})
	})

	group.GET("/torrents", func(ctx *gin.Context) {
		ctx.HTML(200, "torrents.html", gin.H{})
	})

	group.POST("/home", func(c *gin.Context) {
		//get all id from query
		var links []string
		ids := c.PostFormArray("id")
		log.Println(ids)
		for _, id := range ids {
			torrent := w.client.GetTorrentInfo(id)
			links = append(links, torrent.Links[0])
		}

		c.HTML(200, "home.html", gin.H{"links": strings.Join(links, "\n")})
	})

	group.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})
}
