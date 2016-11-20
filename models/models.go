package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName     string
	PasswordHash string
}

type UserSession struct {
	SessionKey string `gorm:"primary_key"`
	UserID     uint   `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type Post struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Title     string
	Body      string
	Tag       string
}
