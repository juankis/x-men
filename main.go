package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	adn := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/mutant", func(c *gin.Context) {
		if isMutant(adn) {
			c.String(200, "")
		} else {
			c.String(403, "")
		}
	})
	router.Run(":9990")
}
