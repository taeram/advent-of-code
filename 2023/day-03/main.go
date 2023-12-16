package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/taeram/advent-of-code/internal/utils"
)

func main() {
	PartNumberSum := 0
	inputs := utils.ReadInputFile()
	for i, input := range inputs {
		before := ""
		if i > 1 {
			before = inputs[i-1]
		}

		after := ""
		if i < len(inputs)-1 {
			after = inputs[i+1]
		}

		partNumbers := GetPartNumbers(before, input, after)
		for _, partNumber := range partNumbers {
			PartNumberSum += partNumber
		}
	}
	utils.Logger("Part Number sum: %d", PartNumberSum)
}

func GetPartNumbers(before string, current string, after string) []int {
	numbers := []int{}

	// Get the symbol indexes.
	symbolIndexes := GetSymbolIndexes(before)
	symbolIndexes = append(symbolIndexes, GetSymbolIndexes(current)...)
	symbolIndexes = append(symbolIndexes, GetSymbolIndexes(after)...)

	// Get the number indexes.
	indexes := regexp.MustCompile(`\d+`).FindAllIndex([]byte(current), -1)
	if len(indexes) > 0 {
		for _, index := range indexes {
			for i := index[0] - 1; i < index[1]+1; i++ {
				if slices.Contains(symbolIndexes, i) {
					number, err := strconv.Atoi(fmt.Sprintf("%v", current[index[0]:index[1]]))
					if err != nil {
						panic(err)
					}

					numbers = append(numbers, number)
					break
				}
			}
		}
	}

	return numbers
}

// GetSymbolIndexes return the position any non-dot, non-number character.
func GetSymbolIndexes(input string) []int {
	var validIndexes []int

	indexes := regexp.MustCompile(`[^.\d]`).FindAllIndex([]byte(input), -1)
	if len(indexes) > 0 {
		for _, index := range indexes {
			validIndexes = append(validIndexes, index[0])
		}
	}

	return validIndexes
}
