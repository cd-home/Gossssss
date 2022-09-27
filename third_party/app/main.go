package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.New()
	//gin.SetMode(gin.ReleaseMode)
	engine.GET("/", func(ctx *gin.Context) {
		//log.Println(ctx.Query("id"))
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	_ = engine.SetTrustedProxies(nil)
	_ = engine.Run(":8080")
}
