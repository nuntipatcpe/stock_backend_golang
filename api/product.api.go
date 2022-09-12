package api

import "github.com/gin-gonic/gin"

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.POST("/profuct", func(ctx *gin.Context) {})
	}
}
