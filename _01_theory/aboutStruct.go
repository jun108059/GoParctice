package _01_theory

// 특징 1: class(in Java), object(in Python, JS) 없음
// 특징 2: constructor(in JS) OR __init__(in Python) 없음
// 대신 struct 있고, constructor 실행해줘야 함

// person struct 구조체
type person struct {
	// 함수도 가질 수 있다!
	name    string
	age     int
	favFood []string
}

func AboutStruct() string {
	favFood := []string{"kimchi", "ramen"}

	// value type 명시해주지 않아도 되지만, 가독성 고려해서 작성하는 것을 추천
	// yjPark := person{"yj park", 28, favFood} // 사실 GoLand IDE 표시해줌..

	yjPark := person{name: "yj park", age: 28, favFood: favFood}
	return yjPark.name
}
