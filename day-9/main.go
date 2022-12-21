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
	x, y   int
	isTail bool
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

	knots := make([]knot, 10)
	for i := range knots {
		knots[i] = knot{5000, 5000, false}
	}
	knots[9] = knot{5000, 5000, true}
	matrix[knots[0].x][knots[0].y] = true
	visited := 1

	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")

		direction := move[0]
		step, _ := strconv.Atoi(move[1])
		head := &knots[0]

		for i := step; i > 0; i-- {
			moveHead(direction, &head.x, &head.y)

			for i := 0; i < len(knots)-1; i++ {
				newHead := head
				if i != 0 {
					newHead = &knots[i]
				}
				newTail := &knots[i+1]

				if isTouching(newHead.x, newHead.y, newTail.x, newTail.y) {
					continue
				} else {
					moveTail(newHead, newTail, matrix, &visited)
				}
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

func moveTail(head, tail *knot, matrix [][]bool, visited *int) {
	// Move tail
	if (*head).y == (*tail).y && math.Abs(float64((*head).x-(*tail).x)) == 2 {
		if (*head).x > (*tail).x {
			(*tail).x = (*tail).x + 1
		} else {
			(*tail).x = (*tail).x - 1
		}
	} else if (*head).x == (*tail).x && math.Abs(float64((*head).y-(*tail).y)) == 2 {
		if (*head).y > (*tail).y {
			(*tail).y = (*tail).y + 1
		} else {
			(*tail).y = (*tail).y - 1
		}
	} else if (*head).x != (*tail).x && (*head).y != (*tail).y {
		if (*head).x-(*tail).x > 0 {
			if (*head).y-(*tail).y > 0 {
				(*tail).x = (*tail).x + 1
				(*tail).y = (*tail).y + 1
			} else {
				(*tail).x = (*tail).x + 1
				(*tail).y = (*tail).y - 1
			}
		} else if (*head).x-(*tail).x < 0 {
			if (*head).y-(*tail).y > 0 {
				(*tail).x = (*tail).x - 1
				(*tail).y = (*tail).y + 1
			} else {
				(*tail).x = (*tail).x - 1
				(*tail).y = (*tail).y - 1
			}
		}
	}
	if tail.isTail && matrix[(*tail).x][(*tail).y] == false {
		matrix[(*tail).x][(*tail).y] = true
		*visited++
	}
}
