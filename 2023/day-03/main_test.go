package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbolIndexes(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	symbols := []struct {
		input   string
		indexes []int
	}{
		{
			input:   `...*......`,
			indexes: []int{3},
		},
		{
			input:   `......#...`,
			indexes: []int{6},
		},
		{
			input:   `.....+.58.`,
			indexes: []int{5},
		},
		{
			input:   `...$.*....`,
			indexes: []int{3, 5},
		},
	}

	for _, symbol := range symbols {
		a.Equal(symbol.indexes, GetSymbolIndexes(symbol.input), symbol.input)
	}
}

func TestGetPartNumbers(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	inputs := []struct {
		before  string
		current string
		after   string
		numbers []int
	}{
		{
			before:  ``,
			current: `467..114..`,
			after:   `...*......`,
			numbers: []int{467},
		},
		{
			before:  `...*......`,
			current: `..35..633.`,
			after:   `......#...`,
			numbers: []int{35, 633},
		},
		{
			before:  `......#...`,
			current: `617*......`,
			after:   `.....+.58.`,
			numbers: []int{617},
		},
		{
			before:  `617*......`,
			current: `.....+.58.`,
			after:   `..592.....`,
			numbers: []int{},
		},
		{
			before:  `.....+.58.`,
			current: `..592.....`,
			after:   `......755.`,
			numbers: []int{592},
		},
		{
			before:  `..592.....`,
			current: `......755.`,
			after:   `...$.*....`,
			numbers: []int{755},
		},
		{
			before:  `...$.*....`,
			current: `.664.598..`,
			after:   ``,
			numbers: []int{664, 598},
		},
	}

	numbersSum := 0
	for _, input := range inputs {
		a.Equal(input.numbers, GetPartNumbers(input.before, input.current, input.after))

		for _, number := range input.numbers {
			numbersSum += number
		}
	}
	a.Equal(4361, numbersSum)
}
