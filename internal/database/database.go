package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TOomaAh/RDTHelper/model"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db    *gorm.DB
	cache *cache.Cache
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewDatabase() *Database {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	c := cache.New(1*time.Hour, 5*time.Hour)

	return &Database{
		db,
		c,
	}
}

func (d *Database) Migrate(v interface{}) {
	d.db.AutoMigrate(v)
}

func (d *Database) FindUserByUsername(username string) (*model.User, error) {
	if user, found := d.cache.Get(username); found {
		return user.(*model.User), nil
	}
	var user model.User
	d.db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}
	d.cache.Set(username, &user, cache.DefaultExpiration)
	return &user, nil
}

func (d *Database) FindUserByUserId(userId int64) (*model.User, error) {
	if user, found := d.cache.Get(fmt.Sprintf("%d", userId)); found {
		return user.(*model.User), nil
	}
	var user model.User
	d.db.Where("id = ?", userId).First(&user)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}
	d.cache.Set(fmt.Sprintf("%d", userId), &user, cache.DefaultExpiration)
	return &user, nil
}

func (d *Database) HasUser() bool {
	var count int64
	d.db.Model(&model.User{}).Count(&count)
	return count > 0
}

func (d *Database) CreateUser(user *model.User) error {
	return d.db.Create(user).Error
}

func (d *Database) GetToken(username string) (string, error) {
	if user, found := d.cache.Get(username); found {
		return user.(*model.User).RdtToken, nil
	}

	var user model.User
	d.db.First(&user)
	if user.ID == 0 {
		return "", ErrUserNotFound
	}
	d.cache.Set(username, &user, cache.DefaultExpiration)
	return user.RdtToken, nil
}
