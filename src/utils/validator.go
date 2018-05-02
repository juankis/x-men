package utils

import (
	"errors"
	"regexp"
	"strings"
)

// IsValidDna mi testo
func IsValidDna(dna []string) error {

	matrixSize := len(dna)

	if matrixSize < 4 {
		return errors.New("La matriz debe ser minimo de 4x4")
	}

	dnaString := strings.Join(dna, "")

	// search character distinct to (A, T, C, G)
	matchInvalid, err := regexp.MatchString("(?i)[^ATCG]", dnaString)
	if err != nil {
		return errors.New("error en regerinxp.MatchStg, error: " + err.Error())
	}
	if matchInvalid {
		return errors.New("las letras de los Strings solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN")
	}
	// Is not NxN matrix
	if len(dnaString) != matrixSize*matrixSize {
		return errors.New("la matriz indicada no conrresponde a una matriz NxN")
	}

	return nil
}
