package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	stacks := make([][]string, 9)

	for scanner.Scan() {
		if scanner.Text() == " 1   2   3   4   5   6   7   8   9 " {
			break
		}
		for i, ch := range scanner.Text() {
			if ch == 32 || ch == 91 || ch == 93 {
				// Skip space and square braces
				continue
			} else {
				stacks[i/4] = append(stacks[i/4], string(ch))
			}
		}
	}

	// read empty line
	scanner.Scan()

	firstStacks := make([][]string, 9)
	secondStacks := make([][]string, 9)
	copy(firstStacks, stacks)
	copy(secondStacks, stacks)

	for scanner.Scan() {
		var move, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &move, &from, &to)

		firstStacks = firstPart(move, from, to, firstStacks)
		secondStacks = secondPart(move, from, to, secondStacks)
	}

	fmt.Println("First part: ")

	for _, stack := range firstStacks {
		fmt.Print(stack[0])
	}

	fmt.Println()

	fmt.Println("Second part: ")
	for _, stack := range secondStacks {
		fmt.Print(stack[0])
	}

}

func firstPart(move int, from int, to int, stacks [][]string) [][]string {
	for i := 0; i < move; i++ {
		letter := stacks[from-1][0]
		stacks[from-1] = stacks[from-1][1:]
		stacks[to-1] = append([]string{letter}, stacks[to-1]...)
	}
	return stacks
}

func secondPart(move int, from int, to int, stacks [][]string) [][]string {
	letters := make([]string, move)
	copy(letters, stacks[from-1][:move])
	stacks[from-1] = stacks[from-1][move:]
	stacks[to-1] = append(letters, stacks[to-1]...)

	return stacks
}
