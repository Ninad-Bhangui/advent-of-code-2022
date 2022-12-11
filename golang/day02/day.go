package day01

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Move int
type Outcome int

// Enum representation in go
const (
	ROCK    Move = 1
	PAPER   Move = 2
	SCISSOR Move = 3
)

const (
	LOSE Outcome = 0
	DRAW Outcome = 3
	WIN  Outcome = 6
)

func Calc1(inputFilepath string) int {
	file, err := os.Open(inputFilepath)
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	var totalScore int
	for scanner.Scan() {
		line := scanner.Text()
		elfMove := getMove(strings.Split(line, " ")[0])
		myMove := getMove(strings.Split(line, " ")[1])

		totalScore += calculateWinScore(elfMove, myMove)
	}
	return totalScore
}

func Calc2(inputFilepath string) int {
	file, err := os.Open(inputFilepath)
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	var totalScore int
	for scanner.Scan() {
		line := scanner.Text()
		elfMove := getMove(strings.Split(line, " ")[0])
		myOutcome := getOutcome(strings.Split(line, " ")[1])
		myMove := getMoveFromOutcome(elfMove, myOutcome)
		totalScore += calculateWinScore(elfMove, myMove)
	}
	return totalScore
}

func getMove(moveStr string) Move {
	switch moveStr {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSOR
	default:
		log.Fatal("unknown move")
	}
	return ROCK //will not be executed
}
func getOutcome(outcomeStr string) Outcome {
	switch outcomeStr {
	case "X":
		return LOSE
	case "Y":
		return DRAW
	case "Z":
		return WIN
	default:
		log.Fatal("unknown move")
	}
	return LOSE //will not be executed
}
func calculateWinScore(elfMove, myMove Move) int {
	return int(myMove) + int(calculateOutcome(elfMove, myMove))
}
func calculateOutcome(elfMove, myMove Move) Outcome {
	if elfMove == myMove {
		return DRAW
	} else if myMove == ROCK && elfMove == SCISSOR {
		return WIN
	} else if myMove == PAPER && elfMove == ROCK {
		return WIN
	} else if myMove == SCISSOR && elfMove == PAPER {
		return WIN
	} else {
		return LOSE
	}
}
func getMoveFromOutcome(elfMove Move, outcome Outcome) Move {
	if outcome == DRAW {
		return elfMove
	} else if (outcome == WIN && elfMove == ROCK) || (outcome == LOSE && elfMove == SCISSOR) {
		return PAPER
	} else if (outcome == WIN && elfMove == PAPER) || (outcome == LOSE && elfMove == ROCK) {
		return SCISSOR
	} else {
		return ROCK
	}
}
