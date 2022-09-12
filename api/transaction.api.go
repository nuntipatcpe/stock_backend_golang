package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {

	transactionAPI := router.Group("api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", createTransaction)
	}
}

func getTransaction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "Get transaction"})
}
func createTransaction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "Create transaction"})
}
