package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDigits(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	calibrationValues := []struct {
		input string
		start int
		end   int
	}{
		{input: "1abc2", start: 1, end: 2},
		{input: "pqr3stu8vwx", start: 3, end: 8},
		{input: "a1b2c3d4e5f", start: 1, end: 5},
		{input: "treb7uchet", start: 7, end: 7},
	}

	for _, value := range calibrationValues {
		start, end := FindDigits(value.input)
		a.Equal(value.start, start)
		a.Equal(value.end, end)
	}
}
