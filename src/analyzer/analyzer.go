package analyzer

import (
	"strings"

	"github.com/juankis/x-men/src/config"
)

var secuenceFound int
var sizeMatriz int

// IsMutant search a secuence in array of strings
func IsMutant(adn []string) bool {
	secuenceFound = 0
	sizeMatriz = len(adn)
	if FindInColumnsRowsAndDiagonals(adn) {
		return true
	}
	return false
}

// Valid return valid if is mutant
func Valid() bool {
	return config.CantSequenceMin <= secuenceFound
}

// FindInColumnsRowsAndDiagonals find secuence in rows, columns and diagonals
func FindInColumnsRowsAndDiagonals(adn []string) bool {
	for i, row := range adn {
		if FindInRow(row) || FindInRow(getColumn(i, adn)) || FindInDiagonals(adn, i) {
			return true
		}
	}
	return false
}

//FindInDiagonals find secuence in diagonals
func FindInDiagonals(adn []string, i int) bool {
	var diagonals [4]string
	for j := 0; j <= i; j++ {
		x := (sizeMatriz - 1) - j
		y := (sizeMatriz - 1) - (i - j)
		diagonals[0] += string(adn[j][i-j])
		diagonals[1] += string(adn[x][y])
		diagonals[2] += string(adn[j][y])
		diagonals[3] += string(adn[y][j])
		if i == sizeMatriz-1 {
			diagonals[1] = ""
			diagonals[2] = ""
		}
	}
	if CheckInDiagonals(diagonals) {
		return true
	}

	return false
}

//FindInRow find secuence in string
func FindInRow(row string) bool {
	for _, letter := range config.GeneticCodeLetters {
		fined := strings.Count(row, strings.Repeat(letter, config.CantSequenceRepetition))
		secuenceFound += fined
		if Valid() {
			return true
		}
	}
	return false
}

//CheckInDiagonals find secuence in array of strings
func CheckInDiagonals(diagonals [4]string) bool {
	for _, diagonal := range diagonals {
		if FindInRow(diagonal) {
			return true
		}
	}
	return false
}

//getColumn get a column from adn
func getColumn(columnIndex int, adn []string) string {
	var column string
	for j := range adn {
		column += string(adn[j][columnIndex])
	}
	return column
}
