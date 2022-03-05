package main

import (
	"fmt"
	"github.com/jun108059/learngo/_01_theory"
	"github.com/jun108059/learngo/_02_accounts"
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

	/*
		3. Functions part
	*/
	// 3-1. argument & return 명시
	fmt.Println(_01_theory.Multiply(2, 2))
	// 3-2. multiple return
	fmt.Println(_01_theory.LenAndUpper("hello"))
	totalLength, upperName := _01_theory.LenAndUpper(fullName)
	fmt.Println(totalLength, upperName)
	// 3-3. 같은 타입 Argument 여러개 받기
	_01_theory.RepeatMe("youngjun", "park", "golang", "practice")
	// 3-4. Naked Return
	fmt.Println(_01_theory.LenAndUpperByNaked("hello"))
	// 3-5. defer
	response := _01_theory.WhatIsDefer("Defer 테스트")
	fmt.Println(response)

	/*
		4. for loop
	*/
	// 4-1. for & range
	total := _01_theory.SuperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	/*
		5. if else
	*/
	fmt.Println(_01_theory.CanIDrink(18))
	fmt.Println(_01_theory.CanIDrinkInKorea(18))

	/*
		6. Switch
	*/
	fmt.Println(_01_theory.CheckGrade(85))
	fmt.Println(_01_theory.CheckRank(100))

	/*
		7. Pointer
	*/
	fmt.Println(_01_theory.Pointer(5300))

	/*
		8. Array And Slice
	*/
	fmt.Println(_01_theory.AboutArray())
	fmt.Println(_01_theory.AboutSlice())

	/*
		9. Map
	*/
	fmt.Println(_01_theory.AboutMap("yj park", "28"))
	fmt.Println(_01_theory.AboutMapForRange("yj park", "28"))

	/*
		10. Struct
	*/
	fmt.Println(_01_theory.AboutStruct())

	/*
		11. accounts (생성자)
	*/
	account := _02_accounts.NewAccount("yj park")
	fmt.Println(account)

	fmt.Println("===============================")

	/*
		12. method (함수)
	*/
	account.Deposit(100)
	fmt.Println(account.GetBalance())
}
