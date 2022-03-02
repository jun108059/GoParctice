package _01_theory

import "fmt"

// AboutMap map 선언 및 구조
func AboutMap(name string, age string) map[string]string {
	// 단점 - value type 강제
	// 개선 - struct
	myProfile := map[string]string{
		"name": name,
		"age":  age}
	return myProfile
}

// AboutMapForRange map 순회 방법
func AboutMapForRange(name string, age string) map[string]string {

	myProfile := map[string]string{
		"name": name,
		"age":  age}

	for key, value := range myProfile {
		fmt.Println(key, value)
	}

	return myProfile
}
