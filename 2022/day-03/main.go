package main

import (
	"strings"

	"github.com/taeram/advent-of-code/internal/utils"
)

const (
	priorityLowercase = 1
	priorityUppercase = 27
)

func main() {
	inputs := utils.ReadInputFile()

	priority := 0
	for _, contents := range inputs {
		priority += getRucksackPriority(contents)
	}
	utils.Logger("Total Priority: %d", priority)

	groupPriority := 0
	for i := 0; i < len(inputs); i += 3 {
		if len(inputs[i]) == 0 {
			continue
		}

		rucksacks := []string{
			inputs[i],
			inputs[i+1],
			inputs[i+2],
		}
		elfGroup := getElfGroup(rucksacks)
		groupPriority += getPriority(elfGroup)
	}
	utils.Logger("Group Priority: %d", groupPriority)
}

func getRucksackPriority(contents string) int {
	contentsList := strings.Split(contents, "")

	contentsLength := len(contents)
	rucksackFirst := contentsList[0 : contentsLength/2]
	rucksackSecond := contentsList[contentsLength/2 : contentsLength]

	priority := 0
	for i := 0; i < len(rucksackSecond); i++ {
		item := rucksackSecond[i]
		if utils.InArray(item, rucksackFirst) {
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

func getElfGroup(rucksacks []string) string {
	alphaList := strings.Split(rucksacks[0], "")
	betaList := strings.Split(rucksacks[1], "")
	deltaList := strings.Split(rucksacks[2], "")

	var badgeItem string
	for i := 0; i < len(deltaList); i++ {
		item := deltaList[i]
		if utils.InArray(item, alphaList) && utils.InArray(item, betaList) {
			badgeItem = item
			break
		}
	}
	return badgeItem
}
