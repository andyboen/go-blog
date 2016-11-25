package main

import (
	"boen/database"
	"boen/models"

	"gopkg.in/gin-gonic/gin.v1"
)

var db, err = database.InitDB()

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	db.DB()
	db.AutoMigrate(&models.Post{}, &models.User{}, &models.UserSession{})

	router.GET("/", index)
	router.GET("/about", about)
	router.POST("/login", login)

	posts := router.Group("/posts")
	{
		posts.GET("/", allPosts)
		posts.GET("/:slug", postByTag)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/posts", index)
		admin.GET("/posts/new", new)
		admin.POST("/posts/create", create)
		admin.GET("/posts/edit/:slug", edit)
		admin.POST("/posts/update", update)
		admin.POST("/posts/delete", delete)
		admin.POST("/posts/undelete", undelete)
	}

	router.Run()
}
