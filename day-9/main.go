package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type knot struct {
	x, y int
}

func main() {
	f, _ := os.Open("/Users/ipekkarakurt/Documents/advent-of-code-2022/day-9/input.txt")
	scanner := bufio.NewScanner(f)

	// matrix of bools, bool denoting whether tail has visited that cell
	// nice magic number
	matrix := make([][]bool, 10001)

	for i := 0; i < 10001; i++ {
		matrix[i] = make([]bool, 10001)
	}

	head := knot{5000, 5000}
	tail := knot{5000, 5000}
	matrix[head.x][head.y] = true
	visited := 1

	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")

		direction := move[0]
		step, _ := strconv.Atoi(move[1])

		for i := step; i > 0; i-- {
			moveHead(direction, &head.x, &head.y)

			if isTouching(head.x, head.y, tail.x, tail.y) {
				continue
			} else {
				moveTail(&head.x, &head.y, &tail.x, &tail.y, matrix, &visited)
			}
		}
	}
	fmt.Println("Visited total: ", visited)
}

func isTouching(headX, headY, tailX, tailY int) bool {
	if math.Abs(float64(headX-tailX)) <= 1 && math.Abs(float64(headY-tailY)) <= 1 {
		return true
	}
	return false
}

func moveHead(direction string, headX, headY *int) {
	// Move head
	switch direction {
	case "L":
		*headX = *headX - 1
	case "R":
		*headX = *headX + 1
	case "U":
		*headY = *headY + 1
	case "D":
		*headY = *headY - 1
	}
}

func moveTail(headX, headY, tailX, tailY *int, matrix [][]bool, visited *int) {
	// Move tail
	if *headY == *tailY && math.Abs(float64(*headX-*tailX)) == 2 {
		if *headX > *tailX {
			*tailX = *tailX + 1
		} else {
			*tailX = *tailX - 1
		}
	} else if *headX == *tailX && math.Abs(float64(*headY-*tailY)) == 2 {
		if *headY > *tailY {
			*tailY = *tailY + 1
		} else {
			*tailY = *tailY - 1
		}
	} else if *headX != *tailX && *headY != *tailY {
		if *headX-*tailX > 0 {
			if *headY-*tailY > 0 {
				*tailX = *tailX + 1
				*tailY = *tailY + 1
			} else {
				*tailX = *tailX + 1
				*tailY = *tailY - 1
			}
		} else if *headX-*tailX < 0 {
			if *headY-*tailY > 0 {
				*tailX = *tailX - 1
				*tailY = *tailY + 1
			} else {
				*tailX = *tailX - 1
				*tailY = *tailY - 1
			}
		}
	}
	if matrix[*tailX][*tailY] == false {
		matrix[*tailX][*tailY] = true
		*visited++
	}
}
