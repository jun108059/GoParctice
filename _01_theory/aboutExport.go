package _01_theory

import "fmt"

// sayBye 소문자로 시작하는 함수는 Private
func sayBye() {
	fmt.Println("Bye")
}

// SayHello 대문자로 시작하는 함수는 export (Public)
func SayHello() {
	fmt.Println("Hello")
}
