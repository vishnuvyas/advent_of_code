package main;

import (
	"fmt"
	"os"
	d1 "github.com/vishnuvyas/advent_of_code/aoc2024/day1" // alias  "day1/problem1" as d1
)

func main() {
	fmt.Println("Hello, AOC World")
	if len(os.Args) >= 2  {
		input_file := os.Args[1]
		fmt.Println("Input file is ", os.Args[1])
		d1.Problem1Solution(input_file)
		d1.Problem2Solution(input_file)
	}
}
