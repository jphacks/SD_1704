package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {

	//以下は後で使います

	//err := model.InsertUser(database.GetInstance().DB, "hoge", "mail", "pass")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//err = model.InsertPost(database.GetInstance().DB, "カレーを飲む青年", 1)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//user, err := model.GetUserByUserId(database.GetInstance().DB, 1)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(user)
	//
	//post, err := model.GetPostById(database.GetInstance().DB, 1)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(post)

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func PostViewHandler(c *gin.Context) {
	//GET METHOD
	if c.Request.Method != "POST" {
		c.Status(http.StatusBadRequest)
	}

	c.HTML(http.StatusOK, "post.html", gin.H{})
}

func PostRegisterHandler(c *gin.Context) {
	//POST METHOD
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
	}
}
