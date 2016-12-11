package main

import (
	"boen/models"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func allPosts(c *gin.Context) {
	posts := models.AllPosts(db)
	c.HTML(http.StatusOK,
		"index_post.tmpl", posts)
}

func viewPost(c *gin.Context) {
	post := models.FindPost(db, c.Param("slug"))
	c.HTML(
		http.StatusOK,
		"view_post.tmpl",
		gin.H{
			"title": post.Title,
			"body":  toHTML(post.Body),
		})
}

func createPost(c *gin.Context) {
	post := models.Post{
		Title:    c.PostForm("title"),
		Body:     c.PostForm("body"),
		Markdown: c.PostForm("body"),
		Slug:     c.PostForm("slug"),
	}
	post.CreatePost(db, post)
	//returnPostStatus(c, "Post Created", post)
}

func newPost(c *gin.Context) {
	c.HTML(http.StatusOK, "form_post.tmpl", gin.H{
		"path": "/admin/posts/create",
	})

}

func editPost(c *gin.Context) {
	post := models.FindPost(db, c.Param("slug"))
	c.HTML(http.StatusOK, "form_post.tmpl", gin.H{
		"path":     "/admin/posts/update",
		"title":    post.Title,
		"markdown": post.Markdown,
		"slug":     post.Slug,
		"id":       post.ID,
	})
}

func updatePost(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	post := models.Post{
		ID:       id,
		Title:    c.PostForm("title"),
		Body:     c.PostForm("body"),
		Markdown: c.PostForm("body"),
		Slug:     c.PostForm("slug"),
	}
	post.UpdatePost(db, post)
	//returnPostStatus(c, "Post Updated", post)
}

func deletePost(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	post := models.Post{ID: id}
	post.DeletePost(db)
	//returnPostStatus(c, "Post Deleted", post)
}

func undeletePost(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	post := models.Post{ID: id}
	post.UnDeletePost(db)
	//returnPostStatus(c, "Post Undeleted", post)
}
