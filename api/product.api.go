package api

import "github.com/gin-gonic/gin"

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"result": "Get product"})
}
func createProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"result": "Create Product"})
}
