package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	var firstPrioSum int
	var secondPrioSum int
	var elfGroup [3]string
	index := 0

	for scanner.Scan() {
		contents := scanner.Text()

		firstPrioSum += firstPart(contents)

		elfGroup[index] = contents
		index++
		if index == 3 {
			secondPrioSum += secondPart(elfGroup)
			index = 0
		}
	}

	fmt.Println("First part:", firstPrioSum)
	fmt.Println("Second part:", secondPrioSum)
}

func firstPart(contents string) int {
	firstCompartment := contents[:len(contents)/2]
	secondCompartment := contents[len(contents)/2:]

	for _, content := range firstCompartment {
		if strings.Contains(secondCompartment, string(content)) {
			prio := getPriority(content)
			return prio
		}
	}
	return 0
}

func secondPart(elfGroup [3]string) int {
	for _, content := range elfGroup[0] {
		if strings.Contains(elfGroup[1], string(content)) && strings.Contains(elfGroup[2], string(content)) {
			prio := getPriority(content)
			return prio
		}
	}
	return 0
}

func getPriority(b int32) int {
	var prio int
	if b > 96 && b < 123 {
		prio = int(b - 96)
	} else if b > 64 && b < 91 {
		prio = int(b - 38)
	}
	return prio
}
