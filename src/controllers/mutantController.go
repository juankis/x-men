package controllers

import (
	"github.com/juankis/x-men/src/db"
	"github.com/juankis/x-men/src/models"
)

func SaveDna(dna []string, isMutant bool) string {
	db := db.Connect()
	defer db.Close()
	mutant := &models.Mutant{Dna: dna, Mutant: isMutant}
	err := db.Insert(mutant)
	if err != nil {
		return ("Error inserting: " + err.Error())
	} else {
		return "1"
	}
}

func CantType(tipe bool) int {
	db := db.Connect()
	defer db.Close()
	count, err := db.Model((*models.Mutant)(nil)).Where("mutant = ?", tipe).Count()
	if err != nil {
		return -1
	} else {
		return count
	}
}

func Humans() int {
	return CantType(false)
}

func Mutants() int {
	return CantType(true)
}
