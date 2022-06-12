package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	// 기본 라우터 생성
	router = gin.Default()

	// HttpStatus : Data
	router.GET("/abc", func(c *gin.Context) {
		// c.String(http.StatusOK, "Hello World!")
		c.JSON(http.StatusOK, gin.H{
			"responseData": "Hello",
		})
	})

	// Http Receive Parameter
	router.GET("/:name", func(c *gin.Context) {
		// :name - gin url 이후에 오는 name 변수로 받음

		var name = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"value": name,
		})
	})

	// 모든 템플릿을 미리 로드 > 다시 로드할 일 없음 > HTML 서빙이 엄청 빠르다!
	router.LoadHTMLGlob("templates/*")

	// Routes 초기화
	initializeRoutes()

	// 애플리케이션 시작!
	router.Run("localhost:8080")

}
