package mathskills

func Median(numbers []float64) float64 {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j && numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	if len(numbers)%2 == 0 {
		return (numbers[len(numbers)/2] + numbers[(len(numbers)/2)-1]) / 2
	} else {
		return numbers[len(numbers)/2]
	}
}
