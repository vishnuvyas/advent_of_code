package day1

import (
	"bufio"
	"log"
	"strings"
	"os"
	"strconv"
)

func readInputFile(filename string) ([]int, []int) {
	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)
	
	file , err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, "\t", 2)
		num1, err1  := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil ||  err2 != nil {
			log.Fatal("Failed parsing on ", line)
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list1, list2
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

