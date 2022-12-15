package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	var totalFirst int
	var totalSecond int

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		firstRangeStr := strings.Split(ranges[0], "-")
		secondRangeStr := strings.Split(ranges[1], "-")

		firstRange := []int{castToInt(firstRangeStr[0]), castToInt(firstRangeStr[1])}
		secondRange := []int{castToInt(secondRangeStr[0]), castToInt(secondRangeStr[1])}

		totalFirst = firstPart(firstRange, secondRange, totalFirst)
		totalSecond = secondPart(firstRange, secondRange, totalSecond)
	}
	fmt.Println("First Part: ", totalFirst)
	fmt.Println("Second Part: ", totalSecond)
}

func castToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
func firstPart(firstRange []int, secondRange []int, total int) int {
	lower := []int{firstRange[0], secondRange[0]}
	higher := []int{firstRange[1], secondRange[1]}

	if lower[0] >= lower[1] && higher[0] <= higher[1] {
		total++
	} else if lower[1] >= lower[0] && higher[1] <= higher[0] {
		total++
	}
	return total
}

func secondPart(firstRange []int, secondRange []int, total int) int {
	if firstRange[1] < secondRange[0] || firstRange[0] > secondRange[1] {
		// Do nothing
	} else {
		total++
	}
	return total
}
