package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

func about(c *gin.Context) {
	c.String(http.StatusOK, "About")
}
