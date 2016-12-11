package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusOK, "404.tmpl", gin.H{})
}
