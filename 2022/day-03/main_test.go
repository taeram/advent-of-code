package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPriority(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, getPriority("a"))
	a.Equal(2, getPriority("b"))
	a.Equal(3, getPriority("c"))
	a.Equal(26, getPriority("z"))
	a.Equal(27, getPriority("A"))
	a.Equal(28, getPriority("B"))
	a.Equal(29, getPriority("C"))
	a.Equal(52, getPriority("Z"))
}

func TestGetRucksackPriority(t *testing.T) {
	a := assert.New(t)

	a.Equal(getPriority("p"), getRucksackPriority("vJrwpWtwJgWrhcsFMMfFFhFp"))
	a.Equal(getPriority("L"), getRucksackPriority("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
	a.Equal(getPriority("P"), getRucksackPriority("PmmdzqPrVvPwwTWBwg"))
	a.Equal(getPriority("v"), getRucksackPriority("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"))
	a.Equal(getPriority("t"), getRucksackPriority("ttgJtRGJQctTZtZT"))
	a.Equal(getPriority("s"), getRucksackPriority("CrZsJsPPZsGzwwsLwLmpwMDw"))
}
