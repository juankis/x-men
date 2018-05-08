package controllers

import (
	"fmt"
	"testing"

	"github.com/juankis/x-men/src/models"
)

func TestSaveDna(t *testing.T) {
	var cases = []models.CaseDna{
		models.CaseDna{
			Dna:      []string{"AALABCDAX", "ECGCEFGAB", "IJCMCJXMC", "LNIYLVYOD", "ABCDCBNDE", "EFGHECGHF", "IJKMIJCMG", "LNSOLNTCH", "AAXALNXOC"},
			Mutant:   false,
			Response: true,
		},
		models.CaseDna{
			Dna:      []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Mutant:   true,
			Response: true,
		},
		models.CaseDna{
			Dna:      []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJVMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Mutant:   false,
			Response: true,
		},
		models.CaseDna{
			Dna:      []string{"AALABCDAT", "ECGCEFGTB", "IJCMCJTMC", "LNICLTYOD", "ABCDVBNDE", "EFGTECGHF", "IJTMIJCMG", "LTSOLNTTH", "TAXALNXOC"},
			Mutant:   true,
			Response: true,
		},
	}

	for i, testCase := range cases {
		res := SaveDna(testCase.Dna, testCase.Mutant)
		if res != testCase.Response {
			t.Errorf("isMutant was incorrect, got: %t, want: %t. Case: %d", res, testCase.Response, i)
		}
	}
}

func TestCantType(t *testing.T) {
	case1 := true
	case2 := true
	fmt.Printf("error: %v", CantType(case1))
	if CantType(case1) == -1 {
		t.Errorf("CantType was incorrect, result is no a number")
	}

	if CantType(case2) == -1 {
		t.Errorf("CantType was incorrect, result is no a number")
	}
}
func TestHumans(t *testing.T) {
	if Humans() == -1 {
		t.Errorf("Humans was incorrect, result is no a number")
	}
}
func TestMutants(t *testing.T) {
	if Mutants() == -1 {
		t.Errorf("Mutants was incorrect, result is no a number")
	}
}
