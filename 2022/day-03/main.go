package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

const priorityLowercase = 1
const priorityUppercase = 27

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file name")
	}

	// Read in the input file.
	inputs := utils.ReadLines(os.Args[1])

	priority := 0
	for _, contents := range inputs {
		priority += getRucksackPriority(contents)
	}
	fmt.Printf("Total Priority: %d\n", priority)
}

func getRucksackPriority(contents string) int {
	var list []string
	for i := 0; i < len(contents)/2; i++ {
		item := string(contents[i])
		list = append(list, item)
	}

	priority := 0
	for i := len(contents) / 2; i < len(contents); i++ {
		item := string(contents[i])
		if utils.InArray(item, list) {
			priority += getPriority(item)
			break
		}
	}

	return priority
}

func getPriority(item string) int {
	itemInt := int([]rune(item)[0])
	if strings.ToLower(item) == item {
		itemInt = itemInt - int('a') + priorityLowercase
	} else {
		itemInt = itemInt - int('A') + priorityUppercase
	}

	return itemInt
}
