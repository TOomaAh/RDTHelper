package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64
	Username string `gorm:"unique"`
	Password string
	ShowAll  bool
	RdtToken string
}

type UserDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	ShowAll  bool   `json:"show_all"`
	RDTToken string `json:"rdt_token"`
}

func (u *User) ToDTO() *UserDTO {
	return &UserDTO{
		ID:       u.ID,
		Username: u.Username,
		ShowAll:  u.ShowAll,
		RDTToken: u.RdtToken,
	}
}

func New(username string, password string, showAll bool, rdtToken string) *User {
	return &User{
		Username: username,
		Password: password,
		ShowAll:  showAll,
		RdtToken: rdtToken,
	}
}

func GetToken(db *gorm.DB) string {
	user := &User{}
	err := db.First(user).Error
	if err != nil {
		return user.RdtToken
	}
	return ""
}

func CountUsers(db *gorm.DB) int64 {
	var count int64
	db.Model(&User{}).Count(&count)
	return count
}

func (u *User) Create(db *gorm.DB) *User {
	db.Create(u)
	if db.Error != nil {
		return nil
	}
	u.ID = db.RowsAffected
	return u
}

func (u *User) Update(db *gorm.DB) *User {
	db.Save(u)
	if db.Error != nil {
		return nil
	}

	return u
}

func (u *User) FindOne(db *gorm.DB) *User {
	db.First(u)
	if db.Error != nil {
		return nil
	}
	return u
}

func (u *User) FindOneByLogin(db *gorm.DB) *User {
	db.Where("username = ?", u.Username).First(u)
	if db.Error != nil {
		return nil
	}
	return u
}

func FindAllUsers(db *gorm.DB) []*User {
	var users []*User
	db.Find(&users)
	return users
}
