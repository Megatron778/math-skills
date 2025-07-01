package mathskillls

func Average(slice []int) int {
	var result int

	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}

	return result / len(slice)
}
