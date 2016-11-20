package main

import (
	"boen/database"
	"boen/models"

	"gopkg.in/gin-gonic/gin.v1"
)

var db, err = database.InitDB()

func main() {
	router := gin.Default()

	db.DB()
	db.AutoMigrate(&models.Post{}, &models.User{}, &models.UserSession{})

	router.GET("/", index)
	router.GET("/about", about)

	posts := router.Group("/posts")
	{
		posts.GET("/", allPosts)
		posts.GET("/:tag", postByTag)
		posts.POST("/new", new)
		posts.POST("/update", update)
		posts.POST("/delete", delete)
		posts.POST("/undelete", undelete)
	}

	router.POST("/login", login)

	router.Run()
}
