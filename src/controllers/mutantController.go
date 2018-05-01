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
