package elf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSection(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	sections := ElfSections{}
	a.Equal(ElfSections{}, sections)

	sections.AddSection(2, 4)
	a.Equal(ElfSections{
		HasOverlap: false,
		Sections: []Section{
			{
				Start: 2,
				End:   4,
			},
		},
	}, sections)

	sections.AddSection(6, 8)
	a.Equal(ElfSections{
		HasOverlap: false,
		Sections: []Section{
			{
				Start: 2,
				End:   4,
			},
			{
				Start: 6,
				End:   8,
			},
		},
	}, sections)
}

func TestNumOverlap(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	sections := ElfSections{}
	sections.AddSection(2, 4)
	sections.AddSection(6, 8)
	a.False(HasOverlap(sections.Sections[0], sections.Sections[1]))

	sections = ElfSections{}
	sections.AddSection(2, 8)
	sections.AddSection(3, 7)
	a.True(HasOverlap(sections.Sections[0], sections.Sections[1]))
}
