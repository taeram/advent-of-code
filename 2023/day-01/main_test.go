package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDigits(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	calibrationValues := []struct {
		input  string
		output int
	}{
		{input: "1abc2", output: 12},
		{input: "pqr3stu8vwx", output: 38},
		{input: "a1b2c3d4e5f", output: 15},
		{input: "treb7uchet", output: 77},
	}

	for _, value := range calibrationValues {
		a.Equal(value.output, FindDigits(value.input))
	}
}

func TestReplaceNumberWords(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	calibrationValues := []struct {
		input  string
		output string
		number int
	}{
		{input: "two1nine", output: "219", number: 29},
		{input: "eightwothree", output: "8wo3", number: 83},
		{input: "abcone2threexyz", output: "abc123xyz", number: 13},
		{input: "xtwone3four", output: "x2ne34", number: 24},
		{input: "4nineeightseven2", output: "4nineeightseven2", number: 42},
		{input: "zoneight234", output: "z1ight234", number: 14},
		{input: "7pqrstsixteen", output: "7pqrst6teen", number: 76},
		{input: "ggdone3nbmsthreefourninefiveoneightpr", output: "ggd13nbms3495on8pr", number: 18},
	}

	for _, value := range calibrationValues {
		a.Equal(value.output, ReplaceNumberWords(value.input))
		a.Equal(value.number, FindDigits(ReplaceNumberWords(value.input)))
	}
}
