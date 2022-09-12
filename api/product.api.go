package api

import (
	"stock/interceptor"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"result": "Get product", "username": ctx.GetString("jwt_username")})
}
func createProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"result": "Create Product"})
}
