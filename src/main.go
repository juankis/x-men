package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/x-men/src/controllers"
	"github.com/juankis/x-men/src/dna"
	"github.com/juankis/x-men/src/dna/utils"
	"github.com/juankis/x-men/src/models"
)

func main() {

	router := gin.Default()

	router.POST("/mutant", handleMutant)
	router.GET("/stats", handleStats)

	router.Run(":8080")

}

func handleMutant(c *gin.Context) {

	var mutantRequest models.MutantRequest

	err := c.BindJSON(&mutantRequest)
	if err != nil {
		c.String(400, "400-Bad-Request")
		return
	}

	validationError := utils.IsValidDna(mutantRequest.Dna)
	if validationError != nil {
		c.String(400, validationError.Error())
		return
	}
	var res bool = dna.IsMutant(mutantRequest.Dna)
	controllers.SaveDna(mutantRequest.Dna, res)
	if res {
		c.String(200, "200-OK")
	} else {
		c.String(403, "403-Forbidden")
	}

	return

}

func handleStats(c *gin.Context) {
	humans := controllers.Humans()
	mutants := controllers.Mutants()
	ratio := float64(mutants) / float64(humans)
	c.JSON(200, gin.H{
		"count_human_dna":  humans,
		"count_mutant_dna": mutants,
		"ratio":            ratio,
	})
	return
}
