package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	r.POST("/dialer", CreateDialer)
	r.POST("/mail", SendMail)
}