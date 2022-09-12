package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func myInterceptor(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "1234" {
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
	}
}

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", myInterceptor, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"result": "Get product"})
}
func createProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"result": "Create Product"})
}
