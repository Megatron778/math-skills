package mathskills

func Variance(numbers []float64, Average float64) float64 {

	var result float64
	for _, v := range numbers {
		result += (v - Average) * (v - Average)
	}
    return 	result / float64(len(numbers))
}