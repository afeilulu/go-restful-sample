package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Name  string    `json:"name"`
	Memo  string    `json:"memo"`
	Users []User    `json:"users" gorm:"foreignkey:GroupID"`
}

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	UserName string    `json:"userName"`
	Phone    string    `json:"phone"`
	GroupID  uuid.UUID `gorm:"type:uuid"`
}
