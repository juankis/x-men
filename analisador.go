package main

import (
	"fmt"
	"strings"
)

const cantSequence int = 4

var letrasCodigoGenetico = []string{"A", "T", "C", "G"}
var secuenceFound int = 0

func isMutant(adn []string) bool {
	if buscarEnFilas(adn) {
		return true
	}
	if buscarEnColumnas() {
		return true
	}
	if buscarEnDiagonales() {
		return true
	}
	return false
}

func validar() bool {
	fmt.Printf("cantidad encontrados: %v \n", secuenceFound)
	return cantSequence <= secuenceFound
}

func buscarEnFilas(adn []string) bool {
	for index := 0; index < len(adn); index++ {
		if buscarEnFila(adn[index]) {
			secuenceFound++
		}
		if validar() {
			return true
		}
	}
	return false
}

func buscarEnColumnas() bool {
	return false
}

func buscarEnDiagonales() bool {
	return false
}

func buscarEnFila(fila string) bool {
	for i := 0; i < len(letrasCodigoGenetico); i++ {
		fined := strings.Repeat(letrasCodigoGenetico[i], cantSequence)
		if strings.Contains(fila, fined) {
			fmt.Printf("Encontrado!\n")
			return true

		}
	}
	return false
}
