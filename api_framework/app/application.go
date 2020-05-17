package app

import (
	"github.com/gin-gonic/gin"
	"microservices/api/log"
)


var (
	router *gin.Engine
)

func init(){
	router = gin.Default()
}

func StartApp() {
	log.Info("starting app","step:1","app:StartApp")
	mapUrls()
	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
