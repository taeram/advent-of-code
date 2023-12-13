package main

import (
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

const (
	choiceRock     = 0
	choicePaper    = 1
	choiceScissors = 2
	choiceDraw     = 1
	choiceWin      = 2
	scoreRock      = 1
	scorePaper     = 2
	scoreScissors  = 3
	scoreLose      = 0
	scoreDraw      = 3
	scoreWin       = 6
)

func main() {
	inputs := utils.ReadInputFile()

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
	utils.Logger("Part 01 - Total Score: %d", scorePart01)
	utils.Logger("Part 02 - Total Score: %d", scorePart02)
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
	//nolint: gocritic
	if meChoice == opponentChoice {
		score += scoreDraw
	} else if (meChoice == choiceRock && opponentChoice == choiceScissors) ||
		(meChoice == choiceScissors && opponentChoice == choicePaper) ||
		(meChoice == choicePaper && opponentChoice == choiceRock) {
		score += scoreWin
	} else {
		score += scoreLose
	}

	switch meChoice {
	case choiceRock:
		score += scoreRock
	case choiceScissors:
		score += scoreScissors
	case choicePaper:
		score += scorePaper
	}

	return score
}

// getChoice returns a rock/paper/scissors choice depending on the opponents choice and the win/lose/draw value.
func getChoice(opponentChoiceStr string, winLoseDrawStr string) string {
	opponentChoice := getChoiceInt(opponentChoiceStr)
	winLoseDraw := getChoiceInt(winLoseDrawStr)
	var meChoice int

	//nolint: gocritic
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
		switch opponentChoice {
		case choiceRock:
			meChoice = choiceScissors
		case choiceScissors:
			meChoice = choicePaper
		case choicePaper:
			meChoice = choiceRock
		}
	}

	return string(rune(meChoice) + rune('X'))
}
