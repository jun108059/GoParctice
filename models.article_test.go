package main

import "testing"

// TestCode getAllArticles 함수 - 전체 Article 조회 되는지
func TestGetAllArticles(t *testing.T) {
	aList := getAllArticles()

	// 검증 1)
	// in memory articleList 저장 길이와 함수로 호출한 ArticleList 개수가 같은지
	if len(aList) != len(articleList) {
		t.Fail()
	}

	// 검증 2)
	// 모든 Article 항목을 검사해서 데이터 같은지 비교
	for i, v := range aList {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
