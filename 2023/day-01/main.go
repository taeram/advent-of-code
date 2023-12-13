package main

import (
	"fmt"
	"strconv"

	"github.com/taeram/advent-of-code/internal/utils"
)

func main() {
	CalibrationSum := 0
	inputs := utils.ReadInputFile()
	for _, input := range inputs {
		start, end := FindDigits(input)
		sum, _ := strconv.Atoi(fmt.Sprintf("%d%d", start, end))
		CalibrationSum += sum
	}
	utils.Logger("Calibration sum: %d", CalibrationSum)
}

func FindDigits(input string) (int, int) {
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

	return start, end
}
