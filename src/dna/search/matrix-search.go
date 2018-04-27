package search

import (
	"sync"

	"github.com/juankis/x-men-2/x-men/src/config"
	"github.com/juankis/x-men-2/x-men/src/dna/utils"
)

// FindInRows xxx
func FindInRows(matrix []string, searchWg *sync.WaitGroup, assertCh chan int) {
	defer searchWg.Done()

	for _, row := range matrix {
		assertCh <- utils.AdnMutantCount(row)
	}

}

// FindInColumns xxx
func FindInColumns(matrix []string, searchWg *sync.WaitGroup, assertCh chan int) {
	defer searchWg.Done()

	var count int

	for col := 0; col < len(matrix); col++ {
		count = 1
		for row := 1; row < len(matrix); row++ {
			if count == 0 || !utils.EqualStringIgnoreCase(string(matrix[row][col]), string(matrix[row-1][col])) {
				count = 1
			} else {
				count++
				if count == config.MutantRepetitions {
					assertCh <- 1
					count = 0
				}
			}
		}
	}

}

// FindInDiagonals cc
func FindInDiagonals(matrix []string, searchWg *sync.WaitGroup, assertCh chan int) {
	defer searchWg.Done()

	sizeMatriz := len(matrix)

	for i := 0; i < sizeMatriz; i++ {
		var diagonals [4]string
		for j := 0; j <= i; j++ {
			x := (sizeMatriz - 1) - j
			y := (sizeMatriz - 1) - (i - j)
			diagonals[0] += string(matrix[j][i-j])
			diagonals[1] += string(matrix[x][y])
			diagonals[2] += string(matrix[j][y])
			diagonals[3] += string(matrix[y][j])
			// for dont duplicate main and minor diagonals
			if i == sizeMatriz-1 {
				diagonals[1] = ""
				diagonals[2] = ""
			}

		}
		for _, diagonal := range diagonals {
			if len(diagonal) < config.MutantRepetitions {
				continue
			}
			assertCh <- utils.AdnMutantCount(diagonal)

		}

	}

}
