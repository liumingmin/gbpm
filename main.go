package main

import (
	"os"
	"time"

	_ "github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	_ "github.com/liumingmin/gbpm/models"
)

func main() {
	if len(os.Args) < 2 {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(os.Args[1])
	}

	port := "8000"
	if len(os.Args) > 2 {
		port = os.Args[2]
	}

	time.Local, _ = time.LoadLocation("Asia/Shanghai")

	router := gin.New()

	if gin.IsDebugging() {
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())

	//controllers.RegisterRouter(router)

	router.Run(":" + port)
}
