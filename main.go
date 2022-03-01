package main

import (
	"fmt"
	"github.com/jun108059/learngo/_01_theory"
)

func main() {
	/*
		1. Packages and Imports
	*/
	_01_theory.SayHello() // 대문자로 시작하면 import 가능
	// _01_theory.sayBye() // 소문자로 시작하면 import 불가능

	/*
		2. Variables and Constants
			- Type Safe
	*/
	// 2-1. constant = 상수 (const in js)
	const name string = "youngjun" // 값을 변경할 수 없음
	fmt.Println(name)

	// 2-2. variable = 변수 (let in js)

	// 2-2-1. 함수 내에서 "축약형" 사용 가능 - 타입 추론
	// var fullName string = "youngjun"
	fullName := "youngjun" // 축약형

	fullName = "youngjun park" // 값 변경 가능
	fmt.Println(fullName)

}
