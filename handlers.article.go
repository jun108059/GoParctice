package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	// 기존에 정의한 getAllArticles 함수를 이용하여 article list 가져오기
	articles := getAllArticles()

	// html 템플릿에 Render 하기 위해 Gin Context -> HTML() 메소드 호출
	c.HTML(
		// HTTP status 200 (OK) 설정
		http.StatusOK,
		// index.html 템플릿 사용 설정
		"index.html",
		// 페이지에서 사용될 데이터 passing(전달)
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	// ID 유효성 검사
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Article 존재 여부 검사
		if article, err := getArticleByID(articleID); err == nil {
			// HTML 메소드 호출 (Template 랜더링)
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// article 데이터를 찾을 수 없음
			_ = c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// article ID 유효하지 않음
		c.AbortWithStatus(http.StatusNotFound)
	}
}
