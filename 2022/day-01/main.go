package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/taeram/advent-of-code/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file name")
	}

	// Read in the input file.
	inputs := utils.ReadLines(os.Args[1])

	// Add up all of the calories carried by each elf.
	elfId := 0
	elfCalories := map[int]int{}
	for _, calories := range inputs {
		if len(calories) == 0 {
			elfId += 1
			elfCalories[elfId] = 0
			continue
		}

		caloriesNum, err := strconv.Atoi(calories)
		if err != nil {
			panic(err)
		}

		elfCalories[elfId] += caloriesNum
	}

	// Create a list, sorted by the number of calories, largest first.
	var sortedCalories []int
	for _, calories := range elfCalories {
		sortedCalories = append(sortedCalories, calories)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedCalories)))

	// Find the max calories carried by a single elf.
	fmt.Printf("Max Calories: %d\n", sortedCalories[0])

	// Find the total calories carried by the top three elves.
	var topThreeCalories int
	for i := 0; i < 3; i++ {
		topThreeCalories += sortedCalories[i]
	}
	fmt.Printf("Top Three Calories: %d\n", topThreeCalories)
}
