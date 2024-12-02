package day1

import (
	"fmt"
	"sort"
)

func Problem1Solution(filename string) {
	list1, list2 := ReadInputFile(filename)
	sort.Ints(list1)
	sort.Ints(list2)
	totalDistance := 0
	
	for i := 0 ; i < len(list1) ; i++ {
		distance := list1[i] - list2[i]
		distance = IntAbs(distance)
		totalDistance += distance
	}

	fmt.Println("Total distance is ", totalDistance)
}
