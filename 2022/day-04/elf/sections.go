package elf

type ElfSections struct {
	Sections   []Section
	HasOverlap bool
}

type Section struct {
	End   int
	Start int
}

func (s *ElfSections) AddSection(start int, end int) {
	s.Sections = append(s.Sections, Section{
		Start: start,
		End:   end,
	})

	if len(s.Sections) > 1 {
		s.HasOverlap = HasOverlap(s.Sections[0], s.Sections[1])
	}
}

func HasOverlap(alpha Section, beta Section) bool {
	return alpha.Start <= beta.Start && alpha.End >= beta.End ||
		beta.Start <= alpha.Start && beta.End >= alpha.End
}
