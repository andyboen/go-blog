package main

import (
	"html/template"
	"strconv"

	"github.com/gin-contrib/sessions"

	"gopkg.in/gin-gonic/gin.v1"
)

// func getSession() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		flashes := session.Flashes()
// 		c.Set("flashes", flashes)
// 		c.Next()
// 	}
// }

func idToString(id string) (convID uint64, err error) {
	convertedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return convertedID, err
}

func invalidIDError(c *gin.Context) {
	Flash(c, "ID not found")
	c.Redirect(400, "/404")
}

func returnStatus(c *gin.Context, status int, flash []interface{}) {
	if status != 0 {
		c.JSON(201, gin.H{
			"flash": flash,
		})
	} else {
		c.JSON(400, gin.H{
			"flash": flash,
		})
	}
}

func Flash(c *gin.Context, message string) {
	session := sessions.Default(c)
	session.AddFlash(message)
	session.Save()
}

func toHTML(s string) template.HTML {
	return template.HTML(s)
}
