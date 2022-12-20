package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	matrix := make([][]int, 0)

	rowIndex := 0
	for scanner.Scan() {
		inner := make([]int, 0)
		for _, c := range scanner.Text() {
			val, _ := strconv.Atoi(string(c))
			inner = append(inner, val)
		}
		matrix = append(matrix, [][]int{inner}...)
		rowIndex++
	}

	rowLength := len(matrix)
	colLength := len(matrix[0])

	part1(matrix, rowLength, colLength)
	part2(matrix, rowLength, colLength)
}

func part2(matrix [][]int, rowLength int, colLength int) {
	maxScore := 0
	for outerIndex := 1; outerIndex < rowLength-1; outerIndex++ {
		for innerIndex := 1; innerIndex < colLength-1; innerIndex++ {
			score := calculateScenicScore(matrix, outerIndex, innerIndex, rowLength, colLength)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println("Max scenic score: ", maxScore)
}

func calculateScenicScore(matrix [][]int, row int, col int, rowLength int, colLength int) int {
	visibleToLeft := 0
	visibleToRight := 0
	visibleToTop := 0
	visibleToBottom := 0

	for i := col - 1; i >= 0; i-- {
		if matrix[row][i] < matrix[row][col] {
			visibleToLeft++
		} else {
			visibleToLeft++
			break
		}
	}

	for i := col + 1; i < colLength; i++ {
		if matrix[row][i] < matrix[row][col] {
			visibleToRight++
		} else {
			visibleToRight++
			break
		}
	}

	for i := row - 1; i >= 0; i-- {
		if matrix[i][col] < matrix[row][col] {
			visibleToTop++
		} else {
			visibleToTop++
			break
		}
	}

	for i := row + 1; i < rowLength; i++ {
		if matrix[i][col] < matrix[row][col] {
			visibleToBottom++
		} else {
			visibleToBottom++
			break
		}
	}

	return visibleToRight * visibleToLeft * visibleToTop * visibleToBottom
}

func part1(matrix [][]int, rowLength int, colLength int) {
	// All trees on edges are visible
	visibleTrees := 2*(rowLength-1) + 2*(colLength-1)

	// Find visible trees inside the edge
	for outerIndex := 1; outerIndex < rowLength-1; outerIndex++ {
		for innerIndex := 1; innerIndex < colLength-1; innerIndex++ {
			if isVisibleFromTop(outerIndex, innerIndex, matrix) ||
				isVisibleFromBottom(outerIndex, innerIndex, colLength, matrix) ||
				isVisibleFromLeft(outerIndex, innerIndex, matrix) ||
				isVisibleFromRight(outerIndex, innerIndex, rowLength, matrix) {
				visibleTrees++
			}
		}
	}
	fmt.Println("Visible trees: ", visibleTrees)
}

func isVisibleFromTop(row int, col int, matrix [][]int) bool {
	isVisible := true
	for i := 0; i < col; i++ {
		if matrix[row][i] >= matrix[row][col] {
			isVisible = false
		}
	}
	return isVisible
}

func isVisibleFromBottom(row int, col int, colLength int, matrix [][]int) bool {
	isVisible := true

	for i := col + 1; i < colLength; i++ {
		if matrix[row][i] >= matrix[row][col] {
			isVisible = false
		}
	}
	return isVisible

}

func isVisibleFromLeft(row int, col int, matrix [][]int) bool {
	isVisible := true
	for i := 0; i < row; i++ {
		if matrix[i][col] >= matrix[row][col] {
			isVisible = false
		}
	}
	return isVisible
}

func isVisibleFromRight(row int, col int, rowLength int, matrix [][]int) bool {
	isVisible := true
	for i := row + 1; i < rowLength; i++ {
		if matrix[i][col] >= matrix[row][col] {
			isVisible = false
		}
	}
	return isVisible
}
