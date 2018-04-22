package main

import "fmt"

func main() {
	adn := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	if isMutant(adn) {
		fmt.Printf("Es mutante")
	} else {
		fmt.Printf("No es mutante")
	}
}
