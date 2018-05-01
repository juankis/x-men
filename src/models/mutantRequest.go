package models

// MutantRequest xx
type MutantRequest struct {
	Dna []string `json:"dna" binding:"required"`
}
