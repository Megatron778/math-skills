package mathskillls

import "fmt"

func Median(slice []int) int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if i != j && slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Println(slice)

	if len(slice)%2 == 0 {
		fmt.Println(slice[(len(slice)/2)-1])
		return slice[len(slice)/2] + slice[(len(slice)/2)-1]
	} else {
		return slice[(len(slice) / 2)]
	}
}
