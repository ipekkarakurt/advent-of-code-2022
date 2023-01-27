package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("/Users/ipekkarakurt/Documents/advent-of-code-2022/day-10/input.txt")
	scanner := bufio.NewScanner(f)

	x := 1
	cycle := 0
	signalSum := 0
	crt := make([][]string, 6)

	for i := range crt {
		crt[i] = make([]string, 40)
		for j := range crt[i] {
			crt[i][j] = "."
		}
	}

	targetCycles := make(map[int]bool)
	for i := 20; i <= 220; i += 40 {
		targetCycles[i] = true
	}

	for scanner.Scan() {
		cmd := scanner.Text()

		switch cmd {
		case "noop":
			cycle++
			if targetCycles[cycle] {
				signalSum += x * cycle
			}
			if math.Abs(float64(x-cycle%40)) <= 1 {
				crt[x%6][x%40] = "#"
			}
		default:
			params := strings.Split(cmd, " ")
			val, _ := strconv.Atoi(params[1])

			if val >= 0 {
				for i := 0; i < 2; i++ {
					cycle++
					if targetCycles[cycle] {
						signalSum += x * cycle
					}
					if math.Abs(float64(x-cycle%40)) <= 1 {
						crt[x%6][x%40] = "#"
					}
				}
				x += val
			} else {
				valAbs := int(math.Abs(float64(val)))
				for i := 0; i < 2; i++ {
					cycle++
					if targetCycles[cycle] {
						signalSum += x * cycle
					}
					if math.Abs(float64(x-cycle%40)) <= 1 {
						crt[x%6][x%40] = "#"
					}
				}

				x -= valAbs
			}
		}
	}
	fmt.Println("First part: ", signalSum)
	fmt.Println("Second part: ")
	printCrt(crt)
}

func printCrt(crt [][]string) {
	for i := range crt {
		for j := range crt[i] {
			fmt.Print(crt[i][j])
		}
		fmt.Println()
	}
}
