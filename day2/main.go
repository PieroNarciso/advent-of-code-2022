package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var winExtraValue map[string]int = map[string]int{
	"rock":    1,
	"paper":   2,
	"scissor": 3,
}

var winnerMove map[string]string = map[string]string{
	"rock": "paper",
	"paper": "scissor",
	"scissor": "rock",
}

var losserMove map[string]string = map[string]string{
	"rock": "scissor",
	"paper": "rock",
	"scissor": "paper",
}

var elvMap map[string]string = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissor",
}

var meMap map[string]string = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}


func getPoints(me string, elve string) int {
	if me == elve {
		return 3 + winExtraValue[me]
	} else if me == "rock" && elve == "scissor" {
		return 6 + winExtraValue[me]
	} else if me == "paper" && elve == "rock" {
		return 6 + winExtraValue[me]
	} else if me == "scissor" && elve == "paper" {
		return 6 + winExtraValue[me]
	} 
	return 0 + winExtraValue[me]
}

func getMove(elveMove string, meOutcome string) string {
	outcome := meMap[meOutcome]	
	switch outcome {
	case "lose":
		return losserMove[elveMove]
	case "draw":
		return elveMove
	default:
		return winnerMove[elveMove]
	}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	fileContent := string(content)
	fileContent = strings.TrimSpace(fileContent)
	allPlays := strings.Split(fileContent, "\n")

	totalPoints := 0

	for _, row := range allPlays {
		plays := strings.Split(row, " ")
		first := plays[0]
		second := plays[1]

		moveElve := elvMap[first]
		moveMe := getMove(moveElve, second)

		matchPoints := getPoints(moveMe, moveElve)
		fmt.Printf("elve: %s vs me: %s -- points: %d\n", moveElve, moveMe, matchPoints)
		totalPoints += matchPoints
	}

	fmt.Println(totalPoints)
}
