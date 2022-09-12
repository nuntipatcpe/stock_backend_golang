package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"stock/db"
	"stock/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product" /*interceptor.JwtVerify ,*/, getProduct)
		productAPI.GET("/product/:id" /*interceptor.JwtVerify ,*/, getProductById)
		productAPI.POST("/product", createProduct)
		productAPI.PUT("/product", editProduct)
	}
}

func getProduct(c *gin.Context) {
	var product []model.Product

	// db.GetDb().Find(&product)
	keyword := c.Query("keyword") //string query -> /api/v2/product?keyword=Ar
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDb().Where("name like ?", keyword).Find(&product)
	} else {
		db.GetDb().Find(&product)
	}
	c.JSON(200, product)

}
func getProductById(c *gin.Context) {
	var product model.Product
	db.GetDb().Where("id = ?", c.Param("id")).First(&product)
	c.JSON(200, product)
}
func createProduct(ctx *gin.Context) {
	product := model.Product{}
	product.Name = ctx.PostForm("name")
	product.Stock, _ = strconv.ParseInt(ctx.PostForm("stock"), 10, 64) //base 10 64bit
	product.Price, _ = strconv.ParseFloat(ctx.PostForm("price"), 64)   //64bit
	product.CreatedAt = time.Now()
	db.GetDb().Create(&product)

	image, _ := ctx.FormFile("image")
	saveImage(image, &product, ctx)

	ctx.JSON(200, gin.H{"result": product})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)

		db.GetDb().Model(&product).Update("image", fileName)
	}
}

func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	db.GetDb().Save(&product)
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})

}
