package main

import (
	"main/api"

	"github.com/gin-gonic/gin"
)

func main() {
	//set * Running in "debug" mode. Switch to "release" mode in production
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// router := gin.Default()
	router.Static("/images", "./uploaded/images")

	api.Setup(router)
	router.Run(":8081")
}
