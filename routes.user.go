package main

import (
	"boen/models"

	"gopkg.in/gin-gonic/gin.v1"
)

func login(c *gin.Context) {
	user := models.FindUser(db, c.PostForm("username"))
	if Authorize(c.PostForm("password"), []byte(user.PasswordHash)) == true {
		session := models.CreateSession(db, user.ID)
		c.SetCookie("Authenticated", string(session), 1, "/login", "localhost", true, false)
	}
}
