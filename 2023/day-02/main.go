package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

const (
	numRedCubes   = 12
	numGreenCubes = 13
	numBlueCubes  = 14
)

func main() {
	ValidGameIds := 0
	inputs := utils.ReadInputFile()
	for _, input := range inputs {
		if IsValidBag(input) {
			ValidGameIds += GetGameId(input)
		}
	}
	utils.Logger("Valid Game IDs sum: %d", ValidGameIds)
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

func IsValidBag(input string) bool {
	// Strip off "Game ##: ".
	input = regexp.MustCompile(`^Game \d+:`).ReplaceAllString(input, "")

	isValid := true
	sets := regexp.MustCompile(`;`).Split(input, -1)
	for _, set := range sets {
		set = strings.Trim(set, " ")

		cubes := regexp.MustCompile(`,`).Split(set, -1)
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
				if numCubes > numRedCubes {
					isValid = false
					break
				}
			case "green":
				if numCubes > numGreenCubes {
					isValid = false
					break
				}
			case "blue":
				if numCubes > numBlueCubes {
					isValid = false
					break
				}
			}
		}
	}

	return isValid
}
