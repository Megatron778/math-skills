package main

import (
	"fmt"
	mathskillls "mathskills/functions"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := os.Args

	if len(os.Args) > 2 {
		fmt.Println("invalid input")
	} 

	data , err := os.ReadFile(filename[1])

	if err != nil {
		fmt.Println("Error :" , err)
	}

	slice := strings.Split(string(data), "\n")

	var dataslice []int
	for _, nb := range slice {
		number, _ := strconv.Atoi(nb)

		dataslice = append(dataslice, number) 
	}

	fmt.Println(dataslice)
	Average := mathskillls.Average(dataslice)
	Median := mathskillls.Median(dataslice)
	Variance := mathskillls.Variance(dataslice, Average)
	fmt.Println(Median + Average)
}
