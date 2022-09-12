package api

import "github.com/gin-gonic/gin"

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}

}

func login(ctx *gin.Context) {
	ctx.JSON(401, gin.H{"result": "login"})
}
func register(ctx *gin.Context) {
	ctx.JSON(401, gin.H{"result": "register"})
}
