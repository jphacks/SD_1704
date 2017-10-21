package handler

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/database"
	"github.com/jphacks/SD_1704/stress/model"
)

func RootHandler(c *gin.Context) {

	err := model.InsertUser(database.GetInstance().DB, "hoge", "mail", "pass")
	if err != nil {
		log.Println(err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
