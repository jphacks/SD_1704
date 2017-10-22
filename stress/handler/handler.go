package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/database"
	"github.com/jphacks/SD_1704/stress/model"

	"strconv"

	"fmt"

	"github.com/jphacks/SD_1704/stress/sessions"
)

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func PostViewHandler(c *gin.Context) {
	// 投稿画面
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
		return
	}

	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}

	if sess.Values["id"] == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{})
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

	id, err := model.InsertPost(database.GetInstance().DB, text, userId)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/mypage")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("%d", id))
}

func ShoutRandomHandler(c *gin.Context) {
	post, err := model.GetRandomPost(database.GetInstance().DB)
	if err != nil {
		log.Println(err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("%d", post.ID))
}

func ShoutHandler(c *gin.Context) {
	// 叫ぶ画面
	postId := c.Param("postId")
	pId, err := strconv.ParseInt(postId, 10, 64)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	post, err := model.GetPostById(database.GetInstance().DB, pId)

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

	nickname, _ := c.GetPostForm("nickname")
	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	if nickname == "" || email == "" || password == "" {
		return
	}

	err := model.InsertUser(database.GetInstance().DB, nickname, email, password)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{})
	}

	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}

	user, err := model.GetUserByEmailAndPass(database.GetInstance().DB, email, password)

	sess.Values["id"] = user.ID
	sess.Values["nickname"] = user.Nickname

	if err := sessions.Save(c.Request, c.Writer, sess); err != nil {
		log.Printf("/login: save session failed: %s", err)
		return
	}

	c.HTML(http.StatusCreated, "mypage.html", gin.H{})

}

func RegisterViewHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusBadRequest)
		return
	}

	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}
	if sess.Values["id"] != nil {
		c.HTML(http.StatusOK, "mypage.html", gin.H{})
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

	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}
	if sess.Values["id"] != nil {
		c.HTML(http.StatusOK, "mypage.html", gin.H{})
		return
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

	userId := sess.Values["id"]

	//uId, err := strconv.ParseInt(, 10, 64)
	//if err != nil {
	//	log.Println(err)
	//}

	posts, err := model.GetPostsByUserId(database.GetInstance().DB, userId.(int64))
	if err != nil {
		log.Println(err)
	}

	log.Println(posts)

	c.HTML(http.StatusOK, "mypage.html", gin.H{
		"posts": posts,
	})
}

func LogoutHandler(c *gin.Context) {
	sess, err := sessions.Get(c.Request, "user")
	if err != nil {
		log.Println(err)
	}

	sess.Values["id"] = nil
	sess.Values["nickname"] = nil

	if err := sessions.Save(c.Request, c.Writer, sess); err != nil {
		log.Printf("/login: save session failed: %s", err)
		return
	}

	log.Println("Logout")

	c.Redirect(http.StatusFound, "/login")
	return
}
