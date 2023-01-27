package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation func(int) int
	testVal   int
	ifTrue    int
	ifFalse   int
}

func main() {
	f, _ := os.Open("/Users/ipekkarakurt/Documents/advent-of-code-2022/day-11/input.txt")
	scanner := bufio.NewScanner(f)

	var monkeys []monkey

	// Skip first line for monkey index
	for scanner.Scan() {
		var monki monkey

		// Line describing starting items array
		scanner.Scan()
		for _, item := range strings.Split(scanner.Text()[len("  Starting items: "):], ", ") {
			startingItem, _ := strconv.Atoi(item)
			monki.items = append(monki.items, startingItem)
		}

		// Line describing monkey operation
		scanner.Scan()
		line := scanner.Text()
		operation := line[len("  Operation: new = old "):]
		operator := operation[:1]
		val := operation[2:]

		switch operator {
		case "+":
			monki.operation = func(n int) int {
				if val != "old" {
					valtoInt, _ := strconv.Atoi(val)
					return (n + valtoInt) / 3
				} else {
					return (n + n) / 3
				}
			}

		case "*":
			monki.operation = func(n int) int {
				if val != "old" {
					valtoInt, _ := strconv.Atoi(val)
					return (n * valtoInt) / 3
				} else {
					return (n * n) / 3
				}
			}
		}

		// Line describing test condition
		scanner.Scan()
		testVal, _ := strconv.Atoi(scanner.Text()[len("  Test: divisible by "):])
		monki.testVal = testVal

		// Action for true outcome
		scanner.Scan()
		trueOut, _ := strconv.Atoi(scanner.Text()[len("    If true: throw to monkey "):])
		monki.ifTrue = trueOut

		// Action for false outcome
		scanner.Scan()
		falseOut, _ := strconv.Atoi(scanner.Text()[len("    If false: throw to monkey "):])
		monki.ifFalse = falseOut

		scanner.Scan()

		monkeys = append(monkeys, monki)
	}

	inspections := make([]int, 8)

	for i := 0; i < 20; i++ {
		for index, monk := range monkeys {
			for _, item := range monk.items {
				inspections[index] = inspections[index] + 1
				outcome := monk.operation(item)
				if outcome%monk.testVal == 0 {
					monkeys[monk.ifTrue].items = append(monkeys[monk.ifTrue].items, outcome)
				} else {
					monkeys[monk.ifFalse].items = append(monkeys[monk.ifFalse].items, outcome)
				}
				monkeys[index].items = monkeys[index].items[1:]
			}
		}
	}

	sort.Ints(inspections)
	fmt.Println("First part : ", inspections[7]*inspections[6])
}
