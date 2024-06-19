package main

import (
	"os"
	"time"

	"github.com/TOomaAh/RDTHelper/internal/database"
	"github.com/TOomaAh/RDTHelper/internal/rdthelper"
	"github.com/TOomaAh/RDTHelper/model"
	"github.com/gin-gonic/gin"
)

func setTimeZone() {
	tz := os.Getenv("TZ")
	if tz == "" {
		time.Local = time.UTC
		return
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		time.Local = time.UTC
		return
	}
	time.Local = loc

}

func init() {
	setTimeZone()
}

func main() {
	db := database.NewDatabase()
	db.Migrate(&model.User{})
	r := gin.Default()

	rdthelper.Run(db, r)
}
