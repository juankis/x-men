package models

//Mutant mutant is a structure used to save the ADNs analyzed
type Mutant struct {
	tableName struct{} `sql:"mutant"`
	ID        int      `sql:",pk"`
	Dna       []string
	Mutant    bool `sql:",notnull"`
}
