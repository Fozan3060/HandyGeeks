package models

import (
	"errors"

	"server/pkg/config"
	"server/pkg/utils"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ContactUsData struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Details   string `json:"details"`
	Questions string `json:"questions"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() error {
	existingUser := User{}
	result := db.Where("email = ?", u.Email).First(&existingUser)
	if result.Error == nil {
		return errors.New("User already exists")
	}

	if hashedPassword, err := utils.HashMyPass(u.Password); err != nil {
		return err
	} else {
		u.Password = hashedPassword
		db.Create(&u)
		return nil
	}

}

func (u *User) GetUserByEmail() error {
	result := db.Where("email = ?", u.Email).First(&u)
	return result.Error
}
