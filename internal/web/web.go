package web

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func NewWebGroup(group *gin.RouterGroup) {

	group.GET("/settings", func(ctx *gin.Context) {
		ctx.HTML(200, "settings.html", gin.H{})
	})

	group.GET("/torrents", func(ctx *gin.Context) {
		ctx.HTML(200, "torrents.html", gin.H{})
	})

	group.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})

	group.POST("/home", func(c *gin.Context) {
		//get all id from query
		var links []string
		ids := c.PostFormArray("id")
		for _, id := range ids {
			values := strings.Split(id, ";")
			if len(values) == 2 {
				l := strings.Split(values[1], ",")
				if len(l) != 0 {
					links = append(links, l...)
				}
			}
		}
		c.HTML(200, "home.html", gin.H{"links": strings.Join(links, "\n")})
	})

}
