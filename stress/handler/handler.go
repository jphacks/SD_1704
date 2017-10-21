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
	//新規投稿
	//POST METHOD
	if c.Request.Method != "POST" {
		c.Status(http.StatusBadRequest)
		return
	}

	text, _ := c.GetPostForm("text")

	if text == "" {
		log.Println("empty")
		return
	}

	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
		return
	}

	userId := sess.Values["id"].(int64)

	err = model.InsertPost(database.GetInstance().DB, text, userId)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/post")
		return
	}

	c.HTML(http.StatusCreated, "mypage.html", gin.H{})
}

func ShoutHandler(c *gin.Context) {
	// 叫ぶ画面

	post, err := model.GetRandomPost(database.GetInstance().DB)
	if err != nil {
		log.Println(err)
	}

	userId := post.UserId

	user, err := model.GetUserByUserId(database.GetInstance().DB, userId)
	if err != nil {
		log.Println(err)
	}

	c.HTML(http.StatusOK, "shout.html", gin.H{
		"post":        post,
		"createrName": user.Nickname,
	})
}

func RegisterInsertHandler(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Status(http.StatusBadRequest)
		return
	}

	// TODO: 登録時もセッションを保つ

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

	c.HTML(http.StatusCreated, "mypage.html", gin.H{})

}

func RegisterViewHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
		return
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
	c.HTML(http.StatusCreated, "mypage.html", gin.H{})
}

func LoginViewHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
	}
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func MyPageHandler(c *gin.Context) {
	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}

	if sess.Values["id"] == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "mypage.html", gin.H{})
}
