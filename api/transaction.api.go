package api

import "github.com/gin-gonic/gin"

func SetupTransactionAPI(router *gin.Engine) {

	transactionAPI := router.Group("api/v2")
	{
		transactionAPI.POST("/transaction", func(ctx *gin.Context) {})
	}
}
