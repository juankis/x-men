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
	sizeMatriz = len(adn)
	if findInRows(adn) || findInColumns(adn) || findInDiagonals(adn) {
		return true
	}
	return false
}

func valid() bool {
	return cantSequenceMin <= secuenceFound
}

func findInRows(adn []string) bool {
	for index := 0; index < sizeMatriz; index++ {
		if findInRow(adn[index]) {
			if addFound() {
				return true
			}
		}
	}
	return false
}

func findInColumns(adn []string) bool {
	for i := 0; i < sizeMatriz; i++ {
		column := ""
		for j := 0; j < sizeMatriz; j++ {
			column += string(adn[j][i])
		}
		if findInRow(column) {
			if addFound() {
				return true
			}
		}
	}
	return false
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
	for i := 0; i < len(geneticCodeLetters); i++ {
		fined := strings.Repeat(geneticCodeLetters[i], cantSequenceRepetition)
		if strings.Contains(row, fined) {
			return true
		}
	}
	return false
}

func checkInDiagonals(diagonals [4]string) bool {
	for i := 0; i < len(diagonals); i++ {
		if findInRow(diagonals[i]) {
			if addFound() {
				return true
			}
		}
	}
	return false
}

func addFound() bool {
	secuenceFound++
	return valid()
}
