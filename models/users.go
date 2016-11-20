package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/jinzhu/gorm"
)

func FindUser(db *gorm.DB, username string) User {
	var user User
	db.Where("user_name = ?", username).First(&user)
	return user
}

func CreateSession(db *gorm.DB, userID uint) string {
	key := generateKey()
	session := UserSession{SessionKey: key, UserID: userID}
	db.NewRecord(session)
	db.Create(&session)
	return session.SessionKey
}

func generateKey() string {
	size := 32
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Println(err)
	}
	rs := base64.URLEncoding.EncodeToString(rb)
	return rs
}
