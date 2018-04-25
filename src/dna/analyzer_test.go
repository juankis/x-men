package dna

import "testing"

type Case struct {
	Matrix []string
	Passed bool
}

func TestIsMutant(t *testing.T) {
	var cases = []Case{
		Case{
			Matrix: []string{"AALABCDAX", "ECGCEFGAB", "IJCMCJXMC", "LNIYLVYOD", "ABCDCBNDE", "EFGHECGHF", "IJKMIJCMG", "LNSOLNTCH", "AAXALNXOC"},
			Passed: false,
		},
		Case{
			Matrix: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Passed: true,
		},
		Case{
			Matrix: []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJVMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Passed: false,
		},
		Case{
			Matrix: []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJTMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Passed: true,
		},
	}

	for i, testCase := range cases {
		isMutant := IsMutant(testCase.Matrix)
		if isMutant != testCase.Passed {
			t.Errorf("isMutant was incorrect, got: %t, want: %t. Case: %d", isMutant, testCase.Passed, i)
		}
	}

}
