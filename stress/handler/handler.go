package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ShoutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "shout.html", gin.H{})
}
