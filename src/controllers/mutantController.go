package controllers

import (
	"github.com/juankis/x-men/src/db"
	"github.com/juankis/x-men/src/models"
	"github.com/juankis/x-men/src/utils"
)

//SaveDna save Dna and if it is mutant
func SaveDna(dna []string, isMutant bool) bool {
	db := db.Connect()
	defer db.Close()
	mutant := &models.Mutant{Dna: dna, Mutant: isMutant}
	err := db.Insert(mutant)
	if err != nil {
		utils.WriteInLog("Error inserting: " + err.Error())
		return false
	}
	return true
}

//CantType return number of rows by atrib mutant
func CantType(tipe bool) int {
	db := db.Connect()
	defer db.Close()
	count, err := db.Model((*models.Mutant)(nil)).Where("mutant = ?", tipe).Count()
	if err != nil {
		utils.WriteInLog("Error consulting: " + err.Error())
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
