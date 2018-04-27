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
	if findInColumnsAndRows(adn) || findInDiagonals(adn) {
		return true
	}
	return false
}

func valid() bool {
	return cantSequenceMin <= secuenceFound
}

func findInColumnsAndRows(adn []string) bool {
	var res bool = false
	for i := 0; i < sizeMatriz; i++ {
		res = findInRow(adn[i])
		if res {
			return true
		}
		column := ""
		for j := 0; j < sizeMatriz; j++ {
			column += string(adn[j][i])
		}
		if findInRow(column) {
			res = true
		}
	}
	return res
}

func findInDiagonals(adn []string) bool {
	for i := 0; i < sizeMatriz; i++ {
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
	}
	return false
}

func findInRow(row string) bool {
	for _, letter := range geneticCodeLetters {
		fined := strings.Count(row, strings.Repeat(letter, cantSequenceRepetition))
		if fined > 0 {
			secuenceFound += fined
			if valid() {
				return true
			}
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
