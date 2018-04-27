package dna

import (
	"strings"
)

const cantSequenceRepetition int = 4
const cantSequenceMin int = 2

var geneticCodeLetters = []string{"A", "T", "C", "G"}
var secuenceFound int
var sizeMatriz int

// IsMutant xxx
func IsMutant(adn []string) bool {
	secuenceFound = 0
	sizeMatriz = len(adn)
	if findInColumnsRowsAndDiagonals(adn) {
		return true
	}
	return false
}

func valid() bool {
	return cantSequenceMin <= secuenceFound
}

func findInColumnsRowsAndDiagonals(adn []string) bool {
	for i, row := range adn {
		if findInRow(row) {
			return true
		}
		column := ""
		for j := range adn {
			column += string(adn[j][i])
		}
		if findInRow(column) {
			return true
		}
		if findInDiagonals(adn, i) {
			return true
		}
	}
	return false
}

func findInDiagonals(adn []string, i int) bool {
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
	if checkInDiagonals(diagonals) {
		return true
	}

	return false
}

func findInRow(row string) bool {
	for _, letter := range geneticCodeLetters {
		fined := strings.Count(row, strings.Repeat(letter, cantSequenceRepetition))
		secuenceFound += fined
		if valid() {
			return true
		}
	}
	return false
}

func checkInDiagonals(diagonals [4]string) bool {
	for _, diagonal := range diagonals {
		if findInRow(diagonal) {
			return true
		}
	}
	return false
}
