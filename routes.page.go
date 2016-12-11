package main

import (
	"boen/models"
	"net/http"

	"github.com/gin-contrib/sessions"

	"gopkg.in/gin-gonic/gin.v1"
)

func viewPage(c *gin.Context) {
	page := models.FindPage(db, c.Param("slug"))
	c.HTML(
		http.StatusOK,
		"view_page.tmpl",
		gin.H{
			"title": page.Title,
			"body":  toHTML(page.Body),
		})
}

func createPage(c *gin.Context) {
	page := models.Page{
		Title:    c.PostForm("title"),
		Body:     c.PostForm("body"),
		Markdown: c.PostForm("body"),
		Slug:     c.PostForm("slug"),
	}
	status, message := page.CreatePage(db, page)
	Flash(c, message)
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()
	returnStatus(c, status, flashes)
}

func newPage(c *gin.Context) {
	c.HTML(http.StatusOK, "form_post.tmpl", gin.H{
		"path": "/admin/pages/create",
	})

}

func editPage(c *gin.Context) {
	page := models.FindPage(db, c.Param("slug"))
	c.HTML(http.StatusOK, "form_post.tmpl", gin.H{
		"path":     "/admin/posts/update",
		"title":    page.Title,
		"markdown": page.Markdown,
		"slug":     page.Slug,
		"id":       page.ID,
	})
}

func updatePage(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	page := models.Page{
		ID:       id,
		Title:    c.PostForm("title"),
		Body:     c.PostForm("body"),
		Markdown: c.PostForm("body"),
		Slug:     c.PostForm("slug"),
	}
	status, message := page.UpdatePage(db, page)
	Flash(c, message)
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()
	returnStatus(c, status, flashes)
}

func deletePage(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	page := models.Page{ID: id}
	page.DeletePage(db)
	// returnPostStatus(c, "Page Deleted", page)
}

func undeletePage(c *gin.Context) {
	id, err := idToString(c.PostForm("id"))
	if err != nil {
		invalidIDError(c)
		return
	}

	page := models.Page{ID: id}
	page.UnDeletePage(db)
	// returnPostStatus(c, "Page Undeleted", page)
}
