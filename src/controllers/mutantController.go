package controllers

import (
	"github.com/juankis/x-men/src/db"
	"github.com/juankis/x-men/src/models"
)

//SaveDna save Dna and if it is mutant
func SaveDna(dna []string, isMutant bool) string {
	db := db.Connect()
	defer db.Close()
	mutant := &models.Mutant{Dna: dna, Mutant: isMutant}
	err := db.Insert(mutant)
	if err != nil {
		return ("Error inserting: " + err.Error())
	}
	return "1"
}

//CantType return number of rows by atrib mutant
func CantType(tipe bool) int {
	db := db.Connect()
	defer db.Close()
	count, err := db.Model((*models.Mutant)(nil)).Where("mutant = ?", tipe).Count()
	if err != nil {
		return -1
	}
	return count
}

//Humans returns number of humans in database
func Humans() int {
	return CantType(false)
}

//Mutants returns number of humans in database
func Mutants() int {
	return CantType(true)
}
