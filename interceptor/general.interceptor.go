package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneralInterceptor1 - call this methos to add interceptor
func GeneralInterceptor1(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "1234" {
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
	}
}
