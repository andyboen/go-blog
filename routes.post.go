package main

import (
	"boen/models"
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func allPosts(c *gin.Context) {
	data := models.AllPosts(db)
	c.String(http.StatusOK,
		data.Title)
}

func postByTag(c *gin.Context) {
	data := models.FindPost(db, c.Param("tag"))
	c.String(http.StatusOK,
		data.Title)
}

func new(c *gin.Context) {
	post := models.Post{
		Title: c.PostForm("title"),
		Body:  c.PostForm("body"),
		Tag:   c.PostForm("tag"),
	}
	post.CreatePost(db, post)
	returnPostStatus(c, "Post Created", post)
}

func update(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	post := models.Post{
		ID:    id,
		Title: c.PostForm("title"),
		Body:  c.PostForm("body"),
		Tag:   c.PostForm("tag"),
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
		"status": "Error: Invalid id",
	})
}

func returnPostStatus(c *gin.Context, status string, p models.Post) {
	c.JSON(200, gin.H{
		"status": status,
		"title":  p.Title,
		"body":   p.Body,
		"tag":    p.Tag,
	})
}
