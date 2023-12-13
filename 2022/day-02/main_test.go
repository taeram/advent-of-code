package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChoiceInt(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Equal(getChoiceInt("A"), 0)
	a.Equal(getChoiceInt("B"), 1)
	a.Equal(getChoiceInt("C"), 2)
	a.Equal(getChoiceInt("X"), 0)
	a.Equal(getChoiceInt("Y"), 1)
	a.Equal(getChoiceInt("Z"), 2)
}

func TestGetScore(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Equal(getScore("A", "Y"), scorePaper+scoreWin)
	a.Equal(getScore("B", "X"), scoreRock+scoreLose)
	a.Equal(getScore("C", "Z"), scoreScissors+scoreDraw)
}

func TestGetChoice(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Equal(getChoice("A", "Y"), "X")
	a.Equal(getChoice("B", "X"), "X")
	a.Equal(getChoice("C", "Z"), "X")
	a.Equal(getChoice("A", "X"), "Z")
	a.Equal(getChoice("B", "Y"), "Y")
}
