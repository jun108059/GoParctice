package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// 기본 라우터 생성
	router = gin.Default()

	// 모든 템플릿을 미리 로드 > 다시 로드할 일 없음 > HTML 서빙이 엄청 빠르다!
	router.LoadHTMLGlob("templates/*")

	// index.html 라우트 (이후에 route handler 함수로 분리)
	router.GET("/", func(c *gin.Context) {
		// html 템플릿에 Render 하기 위해 Gin Context -> HTML() 메소드 호출
		c.HTML(
			// HTTP status 200 (OK) 설정
			http.StatusOK,
			// index.html 템플릿 사용 설정
			"index.html",
			// 페이지에서 사용될 데이터 passing(전달) - 제목
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// 애플리케이션 시작!
	router.Run()

}
