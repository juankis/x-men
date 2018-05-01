package models

type Mutant struct {
	tableName struct{} `sql:"mutant"`
	Id        int      `sql:",pk"`
	Dna       []string
	Mutant    bool `sql:",notnull"`
}
