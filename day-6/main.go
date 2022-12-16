package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("/Users/ipekkarakurt/Documents/advent-of-code-2022/day-6/input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	sequence := scanner.Text()

	fmt.Println("First part: ")
	findNonRepeatingSeq(sequence, 4)
	fmt.Println("Second part: ")
	findNonRepeatingSeq(sequence, 14)
}

func findNonRepeatingSeq(sequence string, length int) {
	nonRepeatingSeq := ""

	for i, c := range sequence {
		if strings.Contains(nonRepeatingSeq, string(c)) {
			i := strings.Index(nonRepeatingSeq, string(c))
			nonRepeatingSeq = nonRepeatingSeq[i+1:]
			nonRepeatingSeq += string(c)
		} else if len(nonRepeatingSeq) == length {
			fmt.Println(i)
			break
		} else {
			nonRepeatingSeq += string(c)
			if len(nonRepeatingSeq) == length {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
