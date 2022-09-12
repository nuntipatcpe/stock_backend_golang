package api

import (
	"stock/db"
	"stock/interceptor"
	"stock/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(ctx *gin.Context) {

	var user model.User
	if ctx.ShouldBind(&user) == nil {
		var queryUser model.User

		if err := db.GetDb().First(&queryUser, "username = ?", user.Username).Error; err != nil {
			ctx.JSON(200, gin.H{"result": "nok", "error": err})
		} else if !checkPasswordHash(user.Password, queryUser.Password) {
			ctx.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		} else {

			ctx.JSON(200, gin.H{"result": "ok", "token": interceptor.JwtSign(queryUser)})
		}
	} else {

		ctx.JSON(401, gin.H{"status": "unable to bind data"})
	}

	// ctx.JSON(200, gin.H{"result": "login"})
}
func register(ctx *gin.Context) {
	var user model.User
	if ctx.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreateAt = time.Now()
		if err := db.GetDb().Create(&user).Error; err != nil {
			ctx.JSON(200, gin.H{"result": "nok", "error": err})
		} else {
			ctx.JSON(200, gin.H{"result": "ok", "data": user})
		}
	} else {

		ctx.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
