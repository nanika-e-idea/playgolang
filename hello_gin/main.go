package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Router
	router := gin.Default()
	// URI,hundler
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})
	// Server
	err := router.Run("127.0.0.1:8888")
	if err != nil {
		log.Fatal("サーバーの起動に失敗しました", err)
	}
}
