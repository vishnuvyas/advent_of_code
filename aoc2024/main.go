package main;

import (
	"fmt"
	"os"
	p1 "github.com/vishnuvyas/advent_of_code/aoc2024/day1" // alias  "day1/problem1" as p1
)

func main() {
	fmt.Println("Hello, AOC World")
	if len(os.Args) >= 2  {
		input_file := os.Args[1]
		fmt.Println("Input file is ", os.Args[1])
		p1.Problem1Solution(input_file)
	}
}
