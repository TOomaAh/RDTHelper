package database

import (
	"errors"
	"log"

	"github.com/TOomaAh/RDTHelper/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewDatabase() *Database {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Database{db}
}

func (d *Database) Migrate(v interface{}) {
	d.db.AutoMigrate(v)
}

func (d *Database) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	d.db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}
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

func (d *Database) GetToken() (string, error) {
	var user model.User
	d.db.First(&user)
	if user.ID == 0 {
		return "", ErrUserNotFound
	}
	return user.RdtToken, nil
}
