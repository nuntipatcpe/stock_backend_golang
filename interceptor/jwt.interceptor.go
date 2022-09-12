package interceptor

import (
	"fmt"
	"net/http"
	"stock/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "87654321"

func JwtSign(payload model.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token

}

func JwtVerify(ctx *gin.Context) {
	tokenString := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["username"])
		level := fmt.Sprintf("%v", claims["level"])
		ctx.Set("jwt_staff_id", staffID)
		ctx.Set("jwt_username", username)
		ctx.Set("jwt_level", level)

		ctx.Next()
	} else {
		ctx.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
		ctx.Abort()
	}
}
