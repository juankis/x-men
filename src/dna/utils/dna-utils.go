package utils

import (
	"strings"

	"github.com/juankis/x-men-2/x-men/src/config"
)

// AdnMutantCount xxx
func AdnMutantCount(s string) int {

	var count int

	for _, letter := range config.GeneticCodeLetters {
		count += strings.Count(s, strings.Repeat(letter, config.MutantRepetitions))
	}

	return count

}

// EqualStringIgnoreCase xxx
func EqualStringIgnoreCase(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}
