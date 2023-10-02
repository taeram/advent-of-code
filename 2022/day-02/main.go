package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

const choiceRock = 0
const choicePaper = 1
const choiceScissors = 2
const choiceDraw = 1
const choiceWin = 2
const scoreRock = 1
const scorePaper = 2
const scoreScissors = 3
const scoreLose = 0
const scoreDraw = 3
const scoreWin = 6

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file name")
	}

	// Read in the input file.
	inputs := utils.ReadLines(os.Args[1])

	var scorePart01 int
	var scorePart02 int
	for _, strategy := range inputs {
		if len(strategy) == 0 {
			continue
		}
		choices := strings.Split(strategy, " ")
		opponentChoice := choices[0]
		meChoice := choices[1]

		scorePart01 += getScore(opponentChoice, meChoice)
		scorePart02 += getScore(opponentChoice, getChoice(opponentChoice, meChoice))
	}
	fmt.Printf("Part 01 - Total Score: %d\n", scorePart01)
	fmt.Printf("Part 02 - Total Score: %d\n", scorePart02)
}

// getChoiceInt returns an integer for the specified character.
func getChoiceInt(choice string) int {
	choiceInt := int([]rune(choice)[0])
	if choiceInt >= int('X') {
		choiceInt -= int('X')
	} else {
		choiceInt -= int('A')
	}
	return choiceInt
}

// getScore returns the score for the opponent choice and my choice.
func getScore(opponentChoiceStr string, meChoiceStr string) int {
	opponentChoice := getChoiceInt(opponentChoiceStr)
	meChoice := getChoiceInt(meChoiceStr)

	var score int
	if meChoice == opponentChoice {
		score += scoreDraw
	} else if (meChoice == choiceRock && opponentChoice == choiceScissors) ||
		(meChoice == choiceScissors && opponentChoice == choicePaper) ||
		(meChoice == choicePaper && opponentChoice == choiceRock) {
		score += scoreWin
	} else {
		score += scoreLose
	}

	if meChoice == choiceRock {
		score += scoreRock
	} else if meChoice == choiceScissors {
		score += scoreScissors
	} else if meChoice == choicePaper {
		score += scorePaper
	}

	return score
}

// getChoice returns a rock/paper/scissors choice depending on the opponents choice and the win/lose/draw value.
func getChoice(opponentChoiceStr string, winLoseDrawStr string) string {
	opponentChoice := getChoiceInt(opponentChoiceStr)
	winLoseDraw := getChoiceInt(winLoseDrawStr)
	var meChoice int

	if winLoseDraw == choiceDraw {
		meChoice = opponentChoice
	} else if winLoseDraw == choiceWin {
		if opponentChoice == choiceRock {
			meChoice = choicePaper
		} else if opponentChoice == choicePaper {
			meChoice = choiceScissors
		} else if opponentChoice == choiceScissors {
			meChoice = choiceRock
		}
	} else {
		if opponentChoice == choiceRock {
			meChoice = choiceScissors
		} else if opponentChoice == choicePaper {
			meChoice = choiceRock
		} else if opponentChoice == choiceScissors {
			meChoice = choicePaper
		}
	}

	return string(rune(meChoice) + rune('X'))
}
