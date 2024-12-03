package day1

import (
	"fmt"
	"slices"
)

func Problem2Solution(filename string) {
	list1, list2 := readInputFile(filename)
	hist := make(map[int]int)

	for i := 0; i < len(list2); i++ {
		hist[list2[i]] += 1
	}

	slices.Sort(list1)
	similarityScore := 0
	for i := 0; i < len(list1); i++ {
		individualScore := list1[i] * hist[list1[i]]
		similarityScore += individualScore
	}

	fmt.Println("The similarity score is ", similarityScore)
}
