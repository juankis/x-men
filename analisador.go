package main

import (
	"strings"
)

const cantSequence int = 4

var letrasCodigoGenetico = []string{"A", "T", "C", "G"}
var secuenceFound int = 0

func isMutant(adn []string) bool {
	if buscarEnFilas(adn) {
		return true
	}
	if buscarEnColumnas(adn) {
		return true
	}
	if buscarEnDiagonales(adn) {
		return true
	}
	return false
}

func validar() bool {
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

func buscarEnColumnas(adn []string) bool {
	for i := 0; i < len(adn); i++ {
		column := ""
		for j := 0; j < len(adn); j++ {
			column += string(adn[j][i])
		}
		if buscarEnFila(column) {
			secuenceFound++
		}
		if validar() {
			return true
		}
	}
	return true
}

func buscarEnDiagonales(adn []string) bool {
	n := len(adn)
	for i := 0; i < n; i++ {
		var diagonals [4]string
		for j := 0; j <= i; j++ {
			x := (n - 1) - j
			y := (n - 1) - (i - j)
			diagonals[0] += string(adn[j][i-j])
			diagonals[1] += string(adn[x][y])
			diagonals[2] += string(adn[j][y])
			diagonals[3] += string(adn[y][j])
		}
		if revisarDiagonales(diagonals) {
			return true
		}
	}
	return false
}

func buscarEnFila(fila string) bool {
	for i := 0; i < len(letrasCodigoGenetico); i++ {
		fined := strings.Repeat(letrasCodigoGenetico[i], cantSequence)
		if strings.Contains(fila, fined) {
			return true
		}
	}
	return false
}

func revisarDiagonales(diagonals [4]string) bool {
	for i := 0; i < len(diagonals); i++ {
		if buscarEnFila(diagonals[i]) {
			secuenceFound++
		}
	}
	if validar() {
		return true
	}
	return false
}
