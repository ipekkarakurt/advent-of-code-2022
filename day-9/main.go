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
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	// matrix of bools, bool denoting whether tail has visited that cell
	// nice magic number
	matrix := make([][]bool, 10001)

	for i := 0; i < 10001; i++ {
		matrix[i] = make([]bool, 10001)
	}

	var head *bool
	var tail *bool
	headX := 5000
	headY := 5000
	tailX := headX
	tailY := headY
	// Starting point
	s := &matrix[headX][headY]
	head = s
	tail = s
	matrix[headX][headY] = true
	visited := 1

	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")

		direction := move[0]
		step, _ := strconv.Atoi(move[1])

		for i := step; i > 0; i-- {
			moveHead(direction, &headX, &headY, head, matrix)

			if isTouching(headX, headY, tailX, tailY) {
				continue
			} else {
				moveTail(&headX, &headY, &tailX, &tailY, matrix, tail, &visited)
			}
		}
	}
	fmt.Println("Visited total: ", visited)
	fmt.Println("Tail: ", tail)
}

func isTouching(headX, headY, tailX, tailY int) bool {
	if math.Abs(float64(headX-tailX)) <= 1 && math.Abs(float64(headY-tailY)) <= 1 {
		return true
	}
	return false
}

func moveHead(direction string, headX, headY *int, head *bool, matrix [][]bool) {
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
	head = &matrix[*headX][*headY]
}

func moveTail(headX, headY, tailX, tailY *int, matrix [][]bool, tail *bool, visited *int) {
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
	tail = &matrix[*tailX][*tailY]
	if matrix[*tailX][*tailY] == false {
		matrix[*tailX][*tailY] = true
		*visited++
	}
}
