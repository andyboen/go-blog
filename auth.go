package main

import (
	"boen/models"

	"golang.org/x/crypto/bcrypt"
)

func Authorize(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return false
	}
	return true
}

func CreateAdmin(username string, password []byte) {
	hash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		panic(err)
	}
	user := models.User{UserName: username, PasswordHash: string(hash)}
	db.NewRecord(user)
	db.Create(&user)
}
