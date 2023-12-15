package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputs = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

var testGames = []Game{
	{
		id: 1,
		sets: []Set{
			{
				blue: 3,
				red:  4,
			},
			{
				red:   1,
				green: 2,
				blue:  6,
			},
			{
				green: 2,
			},
		},
	},
	{
		id: 2,
		sets: []Set{
			{
				blue:  1,
				green: 2,
			},
			{
				green: 3,
				blue:  4,
				red:   1,
			},
			{
				green: 1,
				blue:  1,
			},
		},
	},
	{
		id: 3,
		sets: []Set{
			{
				green: 8,
				blue:  6,
				red:   20,
			},
			{
				blue:  5,
				red:   4,
				green: 13,
			},
			{
				green: 5,
				red:   1,
			},
		},
	},
	{
		id: 4,
		sets: []Set{
			{
				green: 1,
				red:   3,
				blue:  6,
			},
			{
				green: 3,
				red:   6,
			},
			{
				green: 3,
				blue:  15,
				red:   14,
			},
		},
	},
	{
		id: 5,
		sets: []Set{
			{
				red:   6,
				blue:  1,
				green: 3,
			},
			{
				blue:  2,
				red:   1,
				green: 2,
			},
		},
	},
}

func TestGetGameId(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	games := []struct {
		input string
		id    int
	}{
		{
			input: testInputs[0],
			id:    1,
		},
		{
			input: testInputs[1],
			id:    2,
		},
		{
			input: testInputs[2],
			id:    3,
		},
		{
			input: testInputs[3],
			id:    4,
		},
		{
			input: testInputs[4],
			id:    5,
		},
	}

	for _, game := range games {
		a.Equal(game.id, GetGameId(game.input))
	}
}

func TestGetGame(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	for i, game := range testGames {
		a.Equal(game, GetGame(testInputs[i]))
	}
}

func TestIsValidGame(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	games := []struct {
		game    Game
		isValid bool
	}{
		{
			game:    testGames[0],
			isValid: true,
		},
		{
			game:    testGames[1],
			isValid: true,
		},
		{
			game:    testGames[2],
			isValid: false,
		},
		{
			game:    testGames[3],
			isValid: false,
		},
		{
			game:    testGames[4],
			isValid: true,
		},
	}

	for _, game := range games {
		a.Equal(game.isValid, IsValidGame(game.game), game.game.id)
	}
}

func TestGetMinimumSet(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	games := []struct {
		game  Game
		red   int
		green int
		blue  int
	}{
		{
			game:  testGames[0],
			red:   4,
			green: 2,
			blue:  6,
		},
		{
			game:  testGames[1],
			red:   1,
			green: 3,
			blue:  4,
		},
		{
			game:  testGames[2],
			red:   20,
			green: 13,
			blue:  6,
		},
		{
			game:  testGames[3],
			red:   14,
			green: 3,
			blue:  15,
		},
		{
			game:  testGames[4],
			red:   6,
			green: 3,
			blue:  2,
		},
	}

	for _, game := range games {
		red, green, blue := GetMinimumSet(game.game)
		a.Equal(game.red, red)
		a.Equal(game.green, green)
		a.Equal(game.blue, blue)
	}
}
