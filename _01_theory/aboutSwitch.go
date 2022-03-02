package _01_theory

// CheckGrade switch-case 패턴 1
func CheckGrade(score int) string {
	// if-else if 반복을 피할 수 있다.
	switch {
	case score > 90:
		return "A"
	case score > 70:
		return "B"
	case score > 50:
		return "C"
	}
	return "D"
}

// CheckRank switch-case 패턴 2
func CheckRank(score int) string {
	// if-else if 반복을 피할 수 있다.
	switch score {
	case 100:
		return "1등"
	case 95:
		return "2등"
	case 90:
		return "3등"
	}
	return "그 외"
}
