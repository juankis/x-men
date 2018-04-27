package dna

import (
	"strings"
	"sync"
)

const cantSequenceRepetition int = 4
const cantSequenceMin int = 2

var geneticCodeLetters = []string{"A", "T", "C", "G"}
var secuenceFound int
var sizeMatriz int

// IsMutant xxx
func IsMutant(dna []string) bool {
	sizeMatriz = len(dna)
	return asyncCheckMutant(dna)
}

func findInColumnsAndRows(dna []string, searchWg *sync.WaitGroup, assertCh chan int) {
	defer searchWg.Done()
	for _, row := range dna {
		findInRow(row, assertCh)
		var column string
		for _, col := range row {
			column += string(col)
		}
		findInRow(column, assertCh)
	}
}

func findInDiagonals(dna []string, searchWg *sync.WaitGroup, assertCh chan int) {
	defer searchWg.Done()
	for i := range dna {
		var diagonals [4]string
		for j := 0; j <= i; j++ {
			x := (sizeMatriz - 1) - j
			y := (sizeMatriz - 1) - (i - j)
			diagonals[0] += string(dna[j][i-j])
			diagonals[1] += string(dna[x][y])
			diagonals[2] += string(dna[j][y])
			diagonals[3] += string(dna[y][j])

			if i == sizeMatriz-1 {
				diagonals[1] = ""
				diagonals[2] = ""
			}
		}
		checkInDiagonals(diagonals, assertCh)
	}
}

func findInRow(row string, assertCh chan int) {
	for _, letter := range geneticCodeLetters {
		assertCh <- strings.Count(row, strings.Repeat(letter, cantSequenceRepetition))
	}
}

func checkInDiagonals(diagonals [4]string, assertCh chan int) {
	for _, diagonal := range diagonals {
		findInRow(diagonal, assertCh)
	}
}

func asyncCheckMutant(dnaMatrix []string) bool {
	var searchWg sync.WaitGroup
	searchWg.Add(2)
	assertCh := make(chan int, 1)

	go findInColumnsAndRows(dnaMatrix, &searchWg, assertCh)
	go findInDiagonals(dnaMatrix, &searchWg, assertCh)

	go func(searchWg *sync.WaitGroup, assertCh chan int) {
		searchWg.Wait()
		close(assertCh)
	}(&searchWg, assertCh)

	var assert int
	for r := range assertCh {
		assert += r
		if assert >= cantSequenceMin {
			return true
		}
	}
	return false
}
