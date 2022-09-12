package api

import (
	"fmt"
	"os"
	model "stock/Model"
	"stock/interceptor"
	"strconv"

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
	// ctx.JSON(200, gin.H{"result": "Create Product"})

	product := model.Product{}
	product.Name = ctx.PostForm("name")
	product.Stock, _ = strconv.ParseInt(ctx.PostForm("stock"), 10, 64) //base 10 64bit
	product.Price, _ = strconv.ParseFloat(ctx.PostForm("price"), 64)   //64bit
	image, _ := ctx.FormFile("image")
	product.Image = image.Filename

	// extension := filepath.Ext(image.Filename)
	// fileName := fmt.Sprintf("%d%s", product.ID, extension)
	runningDir, _ := os.Getwd()
	filepath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, image.Filename)
	ctx.SaveUploadedFile(image, filepath)

	ctx.JSON(200, gin.H{"result": product})
}
