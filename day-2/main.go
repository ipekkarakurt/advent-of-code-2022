package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Win  = 6
	Loss = 0
	Draw = 3
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	var score_1 int
	var score_2 int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		score_1 = firstPart(parts, score_1)
		score_2 = secondPart(parts, score_2)
	}
	fmt.Println("First part: ", score_1)
	fmt.Println("Second part: ", score_2)
}

func firstPart(parts []string, score int) int {
	translateMoves := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors", "X": "Rock", "Y": "Paper", "Z": "Scissors"}

	score += calculateRoundScoreByMoves(translateMoves[parts[0]], translateMoves[parts[1]])
	return score
}

func calculateRoundScoreByMoves(opponentMove string, myMove string) int {
	var roundScore int
	moveScore := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3}

	if opponentMove == myMove {
		roundScore = Draw + moveScore[myMove]
		return roundScore
	} else if myMove == "Rock" && opponentMove == "Paper" {
		return moveScore[myMove]
	} else if myMove == "Paper" && opponentMove == "Scissors" {
		return moveScore[myMove]
	} else if myMove == "Scissors" && opponentMove == "Rock" {
		return moveScore[myMove]
	} else {
		roundScore = Win + moveScore[myMove]
		return roundScore
	}
}

func secondPart(parts []string, score int) int {
	translateMoves := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
	translateOutcomes := map[string]string{"X": "Lose", "Y": "Draw", "Z": "Win"}

	score += calculateRoundScoreByOutcome(translateMoves[parts[0]], translateOutcomes[parts[1]])
	return score
}

func calculateRoundScoreByOutcome(opponentMove string, outcome string) int {
	var myMove string
	var roundScore int
	moveScore := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3}

	if outcome == "Draw" {
		myMove = opponentMove
		roundScore = Draw + moveScore[myMove]
	} else if outcome == "Lose" {
		if opponentMove == "Rock" {
			myMove = "Scissors"
		} else if opponentMove == "Paper" {
			myMove = "Rock"
		} else if opponentMove == "Scissors" {
			myMove = "Paper"
		}
		roundScore = moveScore[myMove]
	} else {
		if opponentMove == "Rock" {
			myMove = "Paper"
		} else if opponentMove == "Paper" {
			myMove = "Scissors"
		} else if opponentMove == "Scissors" {
			myMove = "Rock"
		}
		roundScore = moveScore[myMove] + Win
	}
	return roundScore
}
