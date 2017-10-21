package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/database"
	"github.com/jphacks/SD_1704/stress/model"
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

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func PostViewHandler(c *gin.Context) {
	// 投稿画面
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
		return
	}

	//GET METHOD
	c.HTML(http.StatusOK, "post.html", gin.H{})
}

func PostInsertHandler(c *gin.Context) {
	//POST METHOD
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
		return
	}


}

func ShoutHandler(c *gin.Context) {
	// 叫ぶ画面

	post, err := model.GetRandomPost(database.GetInstance().DB)
	if err != nil {
		log.Println(err)
	}
	log.Println(post)

	c.HTML(http.StatusOK, "shout.html", gin.H{
		"post": post,
	})
}

func RegisterInsertHandler(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Status(http.StatusBadRequest)
		return
	}

	nickname, _ := c.GetPostForm("nickname")
	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	if nickname == "" || email == "" || password == "" {
		return
	}

	err := model.InsertUser(database.GetInstance().DB, nickname, email, password)
	if err != nil {
		log.Println(err)
	}

	c.Redirect(http.StatusCreated, "/mypage")
}

func RegisterViewHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
		return
	}
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
