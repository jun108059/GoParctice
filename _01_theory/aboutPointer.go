package _01_theory

// Pointer 주소값 & 값 (깊은복사 & 얕은복사)
func Pointer(price int) (int, int) {
	myCoin := price

	// 주소값(&) and 실제값(*)
	bithumb := &myCoin
	upbit := myCoin

	myCoin = 8700
	// *bithumb = 8700
	// myCoin -> 8700

	// 8700 5300
	return *bithumb, upbit
}
