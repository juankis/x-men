package analyzer

import (
	"testing"

	"github.com/juankis/x-men/src/models"
)

func TestIsMutant(t *testing.T) {
	var cases = []models.Case{
		models.Case{
			Matrix: []string{"AALABCDAX", "ECGCEFGAB", "IJCMCJXMC", "LNIYLVYOD", "ABCDCBNDE", "EFGHECGHF", "IJKMIJCMG", "LNSOLNTCH", "AAXALNXOC"},
			Passed: false,
		},
		models.Case{
			Matrix: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Passed: true,
		},
		models.Case{
			Matrix: []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJVMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Passed: false,
		},
		models.Case{
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
