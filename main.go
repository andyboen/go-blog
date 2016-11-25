package main

import (
	"boen/database"
	"boen/models"

	"gopkg.in/gin-gonic/gin.v1"
)

var db, err = database.InitDB()

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	db.DB()
	db.AutoMigrate(&models.Post{}, &models.User{}, &models.UserSession{})

	router.GET("/", index)
	router.GET("/about", about)

	posts := router.Group("/posts")
	{
		posts.GET("/", allPosts)
		posts.GET("/:slug", postByTag)

	}

	admin := router.Group("/admin")
	{
		admin.GET("/", allPosts)
		admin.GET("/new", post)
		admin.POST("/new", new)
		admin.GET("/update/:slug", edit)
		admin.POST("/update", update)
		admin.POST("/delete", delete)
		admin.POST("/undelete", undelete)
	}
	router.POST("/login", login)

	router.Run()
}
