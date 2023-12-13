package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeram/advent-of-code/2022/day-04/elf"
)

func TestNewSections(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Equal(elf.ElfSections{
		HasOverlap: false,
		Sections: []elf.Section{
			{
				Start: 2,
				End:   4,
			},
			{
				Start: 6,
				End:   8,
			},
		},
	}, NewSections("2-4,6-8"))

	a.Equal(elf.ElfSections{
		HasOverlap: false,
		Sections: []elf.Section{
			{
				Start: 2,
				End:   5,
			},
			{
				Start: 3,
				End:   7,
			},
		},
	}, NewSections("2-5,3-7"))
}
