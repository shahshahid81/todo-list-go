package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/datatypes"
)

type User struct {
	gorm.Model
	Email       string         `gorm:"size:255;not null;unique" json:"email"`
	Password    string         `gorm:"size:255;not null;" json:"password"`
	FirstName   string         `gorm:"size:20;not null;" json:"firstName"`
	LastName    string         `gorm:"size:20;not null;" json:"lastName"`
	DateOfBirth datatypes.Date `gorm:"not null;" json:"dateOfBirth"`
}

func (u *User) String() string {
	return fmt.Sprintf("Id: %v, Email: %s, Password: %s, FirstName: %s, LastName: %s, DateOfBirth: %v", u.ID, u.Email, u.Password, u.FirstName, u.LastName, u.DateOfBirth)
}

func (u *User) SaveUser(DB *gorm.DB) (*User, error) {

	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(*gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil

}
