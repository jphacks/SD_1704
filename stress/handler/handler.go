package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/database"
	"github.com/jphacks/SD_1704/stress/model"

	"github.com/jphacks/SD_1704/stress/sessions"
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

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ShoutHandler(c *gin.Context) {

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
	}
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func LoginHandler(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Status(http.StatusBadRequest)
		return
	}

	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	log.Printf("email: %s, pass: %s\n", email, password)

	if email == "" || password == "" {
		log.Println("empty")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	isExists, err := model.IsUserExists(database.GetInstance().DB, email, password)
	if err != nil {
		log.Println(err)
	}

	if !isExists {
		log.Println("not exists")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := model.GetUserByEmailAndPass(database.GetInstance().DB, email, password)
	if err != nil {
		log.Println(err)
		return
	}

	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}

	sess.Values["id"] = user.ID
	sess.Values["nickname"] = user.Nickname

	if err := sessions.Save(c.Request, c.Writer, sess); err != nil {
		log.Printf("/login: save session failed: %s", err)
		return
	}

	log.Println("Login")
	c.Redirect(http.StatusFound, "/mypage")
}

func LoginViewHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
	}
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
