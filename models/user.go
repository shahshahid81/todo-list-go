package models

import (
	"fmt"

	"github.com/shahshahid81/todo-list-go/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/datatypes"
)

type User struct {
	BaseModel
	Email       string         `gorm:"size:255;not null;unique" json:"email"`
	Password    string         `gorm:"size:255;not null;" json:"password"`
	FirstName   string         `gorm:"size:20;not null;" json:"firstName"`
	LastName    string         `gorm:"size:20;not null;" json:"lastName"`
	DateOfBirth datatypes.Date `gorm:"not null;" json:"dateOfBirth"`
}

func LoginCheck(db *gorm.DB, email string, password string) (string, error) {
	u := User{}
	var err error
	err = db.Model(User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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
