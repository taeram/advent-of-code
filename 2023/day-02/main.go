package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int
	green int
	blue  int
}

const (
	numRedCubes   = 12
	numGreenCubes = 13
	numBlueCubes  = 14
)

func main() {
	ValidGameIds := 0
	GameSetPowerSum := 0
	inputs := utils.ReadInputFile()
	for _, input := range inputs {
		game := GetGame(input)
		if IsValidGame(game) {
			ValidGameIds += GetGameId(input)
		}

		red, green, blue := GetMinimumSet(game)
		GameSetPowerSum += red * green * blue
	}
	utils.Logger("Valid Game IDs sum: %d", ValidGameIds)
	utils.Logger("Game Set Power sum: %d", GameSetPowerSum)
}

func GetGameId(input string) int {
	matches := regexp.MustCompile(`^Game (\d+)`).FindSubmatch([]byte(input))
	if len(matches) == 0 {
		panic(fmt.Sprintf("Could not extract Game id from: %s", input))
	}

	gameId, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		panic(err)
	}

	return gameId
}

func GetGame(input string) Game {
	game := Game{
		id:   GetGameId(input),
		sets: []Set{},
	}

	// Strip off "Game ##: ".
	input = regexp.MustCompile(`^Game \d+:`).ReplaceAllString(input, "")

	// Extract the sets.
	setInputs := regexp.MustCompile(`;`).Split(input, -1)
	for _, setInput := range setInputs {
		set := Set{}
		setInput = strings.Trim(setInput, " ")

		// Extract the cubes.
		cubes := regexp.MustCompile(`,`).Split(setInput, -1)
		for _, cube := range cubes {
			matches := regexp.MustCompile(`(\d+) (\w+)`).FindSubmatch([]byte(cube))
			if len(matches) == 0 {
				panic(fmt.Sprintf("No cube matches found in: %s", cube))
			}

			numCubes, err := strconv.Atoi(string(matches[1]))
			if err != nil {
				panic(err)
			}

			cubeColor := string(matches[2])
			switch cubeColor {
			case "red":
				set.red = numCubes
			case "green":
				set.green = numCubes
			case "blue":
				set.blue = numCubes
			}
		}

		game.sets = append(game.sets, set)
	}

	return game
}

func IsValidGame(game Game) bool {
	isValid := true
	for _, set := range game.sets {
		if set.red > numRedCubes ||
			set.green > numGreenCubes ||
			set.blue > numBlueCubes {
			isValid = false
			break
		}
	}

	return isValid
}

func GetMinimumSet(game Game) (int, int, int) {
	red := 0
	green := 0
	blue := 0

	for _, set := range game.sets {
		if set.red > red {
			red = set.red
		}
		if set.green > green {
			green = set.green
		}
		if set.blue > blue {
			blue = set.blue
		}
	}

	return red, green, blue
}
