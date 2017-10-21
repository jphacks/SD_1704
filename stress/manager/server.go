package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/jphacks/SD_1704/stress/config"
	"github.com/jphacks/SD_1704/stress/handler"
)

func Init() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.LoadHTMLGlob("templates/js/*.html")

	r.GET("/", handler.RootHandler)
	r.GET("/shout", handler.ShoutHandler)

	r.Run(":" + config.GetInstance().PORT) // listen and serve on 0.0.0.0:8080
}
