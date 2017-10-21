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

	err = model.InsertPost(database.GetInstance().DB, "カレーを飲む青年", 1)
	if err != nil {
		log.Println(err)
	}

	user, err := model.GetUserByUserId(database.GetInstance().DB, 1)
	if err != nil {
		log.Println(err)
	}

	log.Println(user)

	post, err := model.GetPostById(database.GetInstance().DB, 1)
	if err != nil {
		log.Println(err)
	}
	log.Println(post)

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
