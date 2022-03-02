package _01_theory

// SuperAdd for & range 활용한 이터레이터
func SuperAdd(numbers ...int) int {
	total := 0
	// for {index}, {value} := range {array}
	for _, number := range numbers {
		total += number
	}
	return total
}
