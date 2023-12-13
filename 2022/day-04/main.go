package main

import (
	"strconv"
	"strings"

	"github.com/taeram/advent-of-code/2022/day-04/elf"
	"github.com/taeram/advent-of-code/internal/utils"
)

func main() {
	NumFullyContained := 0
	inputs := utils.ReadInputFile()
	for _, input := range inputs {
		sections := NewSections(input)
		if elf.HasOverlap(sections.Sections[0], sections.Sections[1]) {
			NumFullyContained += 1
		}
	}
	utils.Logger("Found %d overlapping sections", NumFullyContained)
}

func NewSections(input string) elf.ElfSections {
	elfSections := elf.ElfSections{}

	sections := strings.Split(input, ",")
	for _, ranges := range sections {
		values := strings.Split(ranges, "-")

		rangeStart, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}

		rangeEnd, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		elfSections.AddSection(rangeStart, rangeEnd)
	}

	return elfSections
}
