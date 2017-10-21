package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/database"
	"github.com/jphacks/SD_1704/stress/model"
)

func RootHandler(c *gin.Context) {

	model.InsertUser(database.GetInstance().DB, "hoge", "mail", "pass")

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
