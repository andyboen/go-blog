package main

import (
	"boen/models"
	"html/template"
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func toHTML(s string) template.HTML {
	return template.HTML(s)
}

func allPosts(c *gin.Context) {
	posts := models.AllPosts(db)
	c.HTML(http.StatusOK,
		"post_index.tmpl", posts)
}

func postByTag(c *gin.Context) {
	post := models.FindPost(db, c.Param("slug"))
	c.HTML(
		http.StatusOK,
		"post.tmpl",
		gin.H{
			"title": post.Title,
			"body":  toHTML(post.Body),
		})
}

func new(c *gin.Context) {
	post := models.Post{
		Title:    c.PostForm("title"),
		Body:     c.PostForm("body"),
		Markdown: c.PostForm("body"),
		Slug:     c.PostForm("tag"),
	}
	post.CreatePost(db, post)
	returnPostStatus(c, "Post Created", post)
}

func post(c *gin.Context) {
	c.HTML(http.StatusOK, "new_post.tmpl", gin.H{})

}

func edit(c *gin.Context) {
	post := models.FindPost(db, c.Param("slug"))
	c.HTML(http.StatusOK, "edit_post.tmpl", gin.H{
		"title":    post.Title,
		"markdown": post.Markdown,
		"slug":     post.Slug,
		"id":       post.ID,
	})
}

func update(c *gin.Context) {
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
	returnPostStatus(c, "Post Updated", post)
}

func delete(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	post := models.Post{ID: id}
	post.DeletePost(db)
	returnPostStatus(c, "Post Deleted", post)
}

func undelete(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	post := models.Post{ID: id}
	post.UnDeletePost(db)
	returnPostStatus(c, "Post Undeleted", post)
}

func idToString(id string) (convID uint64, err error) {
	convertedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return convertedID, err
}

func invalidIDError(c *gin.Context) {
	c.JSON(400, gin.H{
		"status": "Error: Invalid ID",
	})
}

func returnPostStatus(c *gin.Context, status string, p models.Post) {
	c.JSON(200, gin.H{
		"status": status,
		"title":  p.Title,
		"body":   p.Body,
		"slug":   p.Slug,
	})
}
