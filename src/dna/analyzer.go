package dna

import (
	"sync"

	"github.com/juankis/x-men-2/x-men/src/config"
	"github.com/juankis/x-men-2/x-men/src/dna/search"
)

// IsMutant xxx
func IsMutant(dna []string) bool {
	return asyncCheckMutant(dna)
}

func asyncCheckMutant(dnaMatrix []string) bool {

	var searchWg sync.WaitGroup
	searchWg.Add(3)

	assertCh := make(chan int, 1)

	go search.FindInRows(dnaMatrix, &searchWg, assertCh)

	go search.FindInColumns(dnaMatrix, &searchWg, assertCh)

	go search.FindInDiagonals(dnaMatrix, &searchWg, assertCh)

	go func(searchWg *sync.WaitGroup, assertCh chan int) {
		searchWg.Wait()
		close(assertCh)
	}(&searchWg, assertCh)

	var assert int

	for r := range assertCh {
		assert += r
		if assert >= config.MinMutantRepetitions {
			return true
		}
	}

	return false
}
