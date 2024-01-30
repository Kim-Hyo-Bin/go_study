package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Gin 엔진 초기화
	router := gin.Default()

	// 라우팅 설정
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// 웹 서버 시작
	router.Run(":8080")
}
