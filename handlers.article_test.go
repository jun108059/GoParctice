package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestCode
// GET 요청이 UnAuthenticated 유저로 들어오는 경우
// 1) 핸들러는 HTTP 상태 코드 200 OK 응답
// 2) 반환된 HTML 내부에 "Home Page" 텍스트가 포함된 title tag 여부
func TestShowIndexPageUnauthenticated(t *testing.T) {
	// common_test.go 파일에 helper 함수 호출
	r := getRouter(true)
	// 모든 Article List + GinContext-HTML() 호출
	r.GET("/", showIndexPage)
	// Request 생성
	req, _ := http.NewRequest("GET", "/", nil)

	// common_test.go 파일 - helper 함수 호출
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
