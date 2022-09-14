package main

import (
	"fmt"
	"main/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//set * Running in "debug" mode. Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)
	// router := gin.New()

	router := gin.Default()
	router.Static("/images", "./uploaded/images")
	api.Setup(router)

	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Print("No Port In Heroku")
		router.Run()

	} else {
		fmt.Print("Environment Port : " + port)
		router.Run(":" + port)
	}
	// router.Run(":8081")
}
