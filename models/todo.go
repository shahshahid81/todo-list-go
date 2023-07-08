package models

import (
	"fmt"
)

type Todo struct {
	BaseModel
	UserId      uint   `gorm:"not null;" json:"userId"`
	Title       string `gorm:"size:30;not null;unique" json:"title"`
	Description string `gorm:"size:255;not null;" json:"description"`
}

func (u *Todo) String() string {
	return fmt.Sprintf("Id: %v, User Id: %v, Title: %s, Description: %s", u.ID, u.UserId, u.Title, u.Description)
}
