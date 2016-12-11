package main

import (
	"boen/database"
	"boen/models"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"
)

var db, err = database.InitDB()

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.LoadHTMLGlob("templates/**/*")
	router.Static("assets", "./assets")

	db.DB()
	db.AutoMigrate(&models.Post{}, &models.User{}, &models.Page{}, &models.UserSession{})

	posts := router.Group("/posts")
	{
		posts.GET("/", allPosts)
		posts.GET("/:slug", viewPost)
	}

	pages := router.Group("/")
	{
		pages.GET("", index)
		pages.GET("/page/:slug", viewPage)

		router.POST("/login", login)
		router.GET("/404", notFound)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/posts", allPosts)
		admin.GET("/posts/new", newPost)
		admin.POST("/posts/create", createPost)
		admin.GET("/posts/edit/:slug", editPost)
		admin.POST("/posts/update", updatePost)
		admin.POST("/posts/delete", deletePost)
		admin.POST("/posts/undelete", undeletePost)

		admin.GET("/pages/new", newPage)
		admin.POST("/pages/create", createPage)
		admin.GET("/pages/edit/:slug", editPage)
		admin.POST("/pages/update", updatePage)
		admin.POST("/pages/delete", deletePage)
		admin.POST("/pages/undelete", undeletePage)

	}

	router.Run()
}
