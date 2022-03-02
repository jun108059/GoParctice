package _01_theory

import (
	"fmt"
	"strings"
)

// Multiply 두 정수를 곱해주는 함수
// func 로 함수 선언, return type : 맨 뒤 선언 (필수)
// argument 타입 명시 - 같은 타입의 argument 면 줄일 수 있음 (a, b int)
func Multiply(a int, b int) int {
	return a * b
}

// LenAndUpper 문자의 길이, 대문자 변환 반환 (Return 2개)
func LenAndUpper(str string) (int, string) {
	return len(str), strings.ToUpper(str)
}

// RepeatMe string Type Argument 여러개 받을 수 있음
func RepeatMe(words ...string) {
	fmt.Println(words)
}

// LenAndUpperByNaked Return 변수를 명시할 수 있음
func LenAndUpperByNaked(str string) (length int, uppercase string) {
	length = len(str)
	uppercase = strings.ToUpper(str)
	return
}

// WhatIsDefer 함수가 끝난 뒤 defer 가 실행
func WhatIsDefer(event string) string {
	defer fmt.Println("defer 실행")
	fmt.Println("이벤트 [" + event + "]를 받아서 로직 수행")
	return "response!"
}
