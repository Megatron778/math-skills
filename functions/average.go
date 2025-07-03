package mathskills

func Average(numbers []float64) float64 {
	var result float64
	for _, v := range numbers {
		result += v
	}
	return result / float64(len(numbers))
}
