package _01_theory

// AboutArray - 배열의 길이를 미리 할당
func AboutArray() [5]string {
	colors := [5]string{"red", "orange", "pink"}
	colors[3] = "blue"
	colors[4] = "green"
	// colors[5] = "white" // out of bounds
	return colors
}

// AboutSlice - 배열 길이를 동적으로 할당
func AboutSlice() []string {
	colors := []string{"red", "orange", "pink"}
	// append - slice 인자를 추가해서 새로운 Type slice 반환
	colors = append(colors, "blue")

	return colors
}
