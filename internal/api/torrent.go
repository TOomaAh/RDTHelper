package api

import (
	"time"

	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/internal/middleware"
	"github.com/TOomaAh/RDTHelper/model"
	"github.com/TOomaAh/RDTHelper/pkg/realdebrid"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type TorrentApi struct {
	cache *cache.Cache
}

func NewTorrentApi(group *gin.RouterGroup, db *database.Database) {

	t := &TorrentApi{
		cache: cache.New(60*time.Minute, 120*time.Minute),
	}

	group.Use(middleware.CheckAuthenticated(db))

	group.Use(func(c *gin.Context) {
		// get rdtClient and catch error
		if user, exist := c.Get("user"); exist {
			if rd, exist := t.cache.Get("rd-" + user.(*model.User).Username); exist {
				c.Set("rd", rd)
			} else {
				client := realdebrid.NewRealDebridClient(user.(*model.User).RdtToken)
				c.Set("rd", client)
				t.cache.Set("rd-"+user.(*model.User).Username, client, cache.DefaultExpiration)
			}
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		}
	})

	group.GET("/torrents", t.getAll)
	group.GET("/torrents/:id", t.getOne)
	group.GET("/torrents/accept/:id", t.acceptOne)
	group.POST("/torrent/upload", t.upload)
	group.POST("/torrents/debrid", t.Debrid)
	group.DELETE("/torrents/:id", t.deleteOne)

}

func (t *TorrentApi) getAll(c *gin.Context) {
	client := c.MustGet("rd").(*realdebrid.RealDebridClient)
	torrents, err := client.GetTorrents()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, torrents)
	}
	torrents = nil
}

func (t *TorrentApi) getOne(c *gin.Context) {
	client := c.MustGet("rd").(*realdebrid.RealDebridClient)
	torrent := client.GetTorrentInfo("")

	if torrent != nil {
		c.JSON(200, torrent)
	} else {
		c.JSON(404, gin.H{"error": "Torrent not found"})
	}
}

func (t *TorrentApi) upload(c *gin.Context) {
	//Upload torrent
	client := c.MustGet("rd").(*realdebrid.RealDebridClient)

	from, err := c.MultipartForm()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	files := from.File["file"]

	if len(files) == 0 {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	err = client.UploadTorrent(files)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Torrent uploaded"})
	}

}

func (t *TorrentApi) Debrid(c *gin.Context) {
	//Debrid torrent
	client := c.MustGet("rd").(*realdebrid.RealDebridClient)

	var linkRequest realdebrid.LinkRequest

	c.ShouldBindJSON(&linkRequest)

	link, err := client.DebridTorrent(linkRequest.Link)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, link)
	}
}

func (t *TorrentApi) acceptOne(c *gin.Context) {
	//Accept one torrent
	client := c.MustGet("rd").(*realdebrid.RealDebridClient)

	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	client.AcceptTorrent(id)

}

func (t *TorrentApi) deleteOne(c *gin.Context) {
	//Delete one torrent
	client := c.MustGet("rd").(*realdebrid.RealDebridClient)
	err := client.DeleteTorrent("nil")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Torrent deleted"})
	}
}
