package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/config"
	"github.com/jphacks/SD_1704/stress/handler"
)

func Init() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", handler.RootHandler)

	r.GET("/post", handler.PostViewHandler)
	r.POST("/post", handler.PostInsertHandler)

	r.GET("/shout/", handler.ShoutRandomHandler)
	r.GET("/shout/:postId", handler.ShoutHandler)

	r.GET("/register", handler.RegisterViewHandler)
	r.POST("/register", handler.RegisterInsertHandler)

	r.GET("/login", handler.LoginViewHandler)
	r.POST("/login", handler.LoginHandler)

	r.GET("/logout", handler.LogoutHandler)

	r.GET("/mypage", handler.MyPageHandler)

	r.Run(":" + config.GetInstance().PORT) // listen and serve on 0.0.0.0:8080
}
