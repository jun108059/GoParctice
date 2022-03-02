package _01_theory

// CanIDrink if-else 표현
func CanIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

// CanIDrinkInKorea if-else & variable expression
func CanIDrinkInKorea(age int) bool {
	// if 조건문에서 생성된 변수는 조건문에서만 사용된다.
	if koreanAge := age - 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}
