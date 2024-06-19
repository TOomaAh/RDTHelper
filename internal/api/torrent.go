package api

import (
	"github.com/TOomaAh/RDTHelper/internal/middleware"
	"github.com/TOomaAh/RDTHelper/pkg/realdebrid"
	"github.com/gin-gonic/gin"
)

type TorrentApi struct {
	client *realdebrid.RealDebridClient
}

func NewTorrentApi(group *gin.RouterGroup) {

	t := &TorrentApi{}

	group.Use(middleware.CheckAuthenticated)

	group.Use(func(c *gin.Context) {
		// get rdtClient and catch error
		client, exist := c.Get("rd")
		if !exist {
			return
		}
		t.client = client.(*realdebrid.RealDebridClient)
	})

	group.Use(func(c *gin.Context) {
		if t.client == nil {
			c.JSON(401, gin.H{"error": "Please login first"})
			c.Abort()
		}
	})

	group.GET("/torrents", t.getAll)
	group.GET("/torrents/:id", t.getOne)
	group.GET("/torrents/accept/:id", t.acceptOne)
	group.POST("/torrent/upload", t.upload)
	group.POST("/torrents/debrid", t.Debrid)
	group.DELETE("/torrents/:id", t.deleteOne)

}

func (t *TorrentApi) getAll(ctx *gin.Context) {
	torrents, err := t.client.GetTorrents()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, torrents)
	}
	torrents = nil
}

func (t *TorrentApi) getOne(c *gin.Context) {
	torrent := t.client.GetTorrentInfo("")

	if torrent != nil {
		c.JSON(200, torrent)
	} else {
		c.JSON(404, gin.H{"error": "Torrent not found"})
	}
}

func (t *TorrentApi) upload(c *gin.Context) {
	//Upload torrent
	t.client.UploadTorrent()
}

func (t *TorrentApi) Debrid(c *gin.Context) {
	//Debrid torrent
	t.client.DebridTorrent()
}

func (t *TorrentApi) acceptOne(c *gin.Context) {
	//Accept one torrent
	t.client.AcceptTorrent("nil")
}

func (t *TorrentApi) deleteOne(c *gin.Context) {
	//Delete one torrent
	t.client.DeleteTorrent("nil")
}
