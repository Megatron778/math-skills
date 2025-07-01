package mathskillls

func Variance(slice []int,Average int) int {
	var result int
	for i := 0; i < len(slice); i++ {
		result += (slice[i] - Average) * (slice[i] - Average) 
	}
	return result / len(slice)
}