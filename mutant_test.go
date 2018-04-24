package main

import "testing"

type Case struct {
	Matriz   []string
	Response bool
}

func TestIsMutant(t *testing.T) {
	var cases = []Case{
		Case{
			Matriz:   []string{"AALABCDAX", "ECGCEFGAB", "IJCMCJXMC", "LNIYLVYOD", "ABCDCBNDE", "EFGHECGHF", "IJKMIJCMG", "LNSOLNTCH", "AAXALNXOC"},
			Response: false,
		},
		Case{
			Matriz:   []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Response: true,
		},
		Case{
			Matriz:   []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJVMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Response: false,
		},
		Case{
			Matriz:   []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJTMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Response: true,
		},
	}

	for i := 0; i < len(cases); i++ {
		isMutant := isMutant(cases[i].Matriz)
		if isMutant != cases[i].Response {
			t.Errorf("isMutant was incorrect, got: %d, want: %d. Case: %d", isMutant, cases[i].Response, i+1)
		}
	}
}
