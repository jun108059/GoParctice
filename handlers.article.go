package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	// 기존에 정의한 getAllArticles 함수를 이용하여 article list 가져오기
	articles := getAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
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

// render
// c : gin Context
// data : 데이터 passing
// templateName : 랜더링할 템플릿 이름
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
