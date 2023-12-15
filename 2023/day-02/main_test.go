package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGameId(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	games := []struct {
		input string
		id    int
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			id:    1,
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			id:    2,
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			id:    3,
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			id:    4,
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			id:    5,
		},
	}

	for _, game := range games {
		a.Equal(game.id, GetGameId(game.input))
	}
}

func TestIsValidBag(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	games := []struct {
		input   string
		isValid bool
	}{
		{
			input:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			isValid: true,
		},
		{
			input:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			isValid: true,
		},
		{
			input:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			isValid: false,
		},
		{
			input:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			isValid: false,
		},
		{
			input:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			isValid: true,
		},
	}

	for _, game := range games {
		a.Equal(game.isValid, IsValidBag(game.input), game.input)
	}
}
