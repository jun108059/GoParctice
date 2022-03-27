package main

import "errors"

// article struct
type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// article struct를 in memory List에 저장 (실 운영은 DB에 연결해서 가져오자)
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// 전체 article list 조회
func getAllArticles() []article {
	return articleList
}

// getArticleByID ID로 Article 조회
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
