package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
  	"mathskills/functions"
)

func main() {
	Arg := os.Args
	if len(Arg) != 2 {
		fmt.Println("invalid input: Please entre one arguments")
		return
	}

	if !strings.HasSuffix(Arg[1], ".txt") && !strings.HasSuffix(Arg[1], ".TXT") {
		fmt.Println("Invalid input: Please use a .txt extension")
		return
	}

	data, err := os.ReadFile(Arg[1])
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	dataslice := strings.Split(string(data), "\n")

	dataslice = mathskills.CleanSlice(dataslice)
	if len(dataslice) == 0 {
		fmt.Println("Error:", Arg[1], "is empty")
		return
	}
	var datafloat []float64
	for _, v := range dataslice {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error Converting :", err)
			return
		}
		datafloat = append(datafloat, float64(number))
	}

	Average := mathskills.Average(datafloat)
	Median := mathskills.Median(datafloat)
	Variance := mathskills.Variance(datafloat, Average)
	StandardDeviation := math.Sqrt(Variance)

	fmt.Println("Average:", int(math.Round(Average)))
	fmt.Println("Median:", int(math.Round(Median)))
	fmt.Println("Variance:", int(math.Round(Variance)))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDeviation)))
}
