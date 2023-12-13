package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

func main() {
	CalibrationSum := 0
	CalibrationSumSecond := 0
	inputs := utils.ReadInputFile()
	for _, input := range inputs {
		CalibrationSum += FindDigits(input)
		CalibrationSumSecond += FindDigits(ReplaceNumberWords(input))
	}
	utils.Logger("Part 1 calibration sum: %d", CalibrationSum)
	utils.Logger("Part 2 calibration sum: %d", CalibrationSumSecond)
}

func FindDigits(input string) int {
	var start int
	var end int

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			num, _ := strconv.Atoi(string(input[i]))
			if start == 0 {
				start = num
			}
			end = num
		}
	}

	sum, _ := strconv.Atoi(fmt.Sprintf("%d%d", start, end))
	return sum
}

func ReplaceNumberWords(input string) string {
	type Numbers struct {
		word  string
		value string
	}

	numbers := []Numbers{
		{word: "one", value: "1"},
		{word: "two", value: "2"},
		{word: "three", value: "3"},
		{word: "four", value: "4"},
		{word: "five", value: "5"},
		{word: "six", value: "6"},
		{word: "seven", value: "7"},
		{word: "eight", value: "8"},
		{word: "nine", value: "9"},
	}

	// Iterate over the input string, one character at a time.
	// Find the first number word in the input string.
	for i := 0; i < len(input); i++ {
		// If the current character is a digit, stop processing.
		if input[i] >= '0' && input[i] <= '9' {
			break
		}

		for _, number := range numbers {
			// If the input starts with a number at the current position, replace it with the numerical value.
			if strings.HasPrefix(input[i:], number.word) {
				input = strings.Replace(input, number.word, number.value, 1)
				break
			}
		}
	}

	// Find the last number word in the input string.
	for i := len(input) - 1; i >= 0; i-- {
		// If the current character is a digit, stop processing.
		if input[i] >= '0' && input[i] <= '9' {
			break
		}

		for _, number := range numbers {
			// If the input starts with a number at the current position, replace it with the numerical value.
			if strings.HasPrefix(input[i:], number.word) {
				input = strings.Replace(input, number.word, number.value, 1)
				break
			}
		}
	}

	return input
}
