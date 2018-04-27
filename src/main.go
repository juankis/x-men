package main

import (
	"bitbucket.org/juanki1989/x-men/src/dna"
	"github.com/gin-gonic/gin"
)

func main() {
	adn := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	router := gin.Default()

	router.POST("/mutant", func(c *gin.Context) {
		if dna.IsMutant(adn) {
			c.String(200, "")
		} else {
			c.String(403, "")
		}
	})

	router.Run(":9990")

}
