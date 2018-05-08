package db

import "github.com/go-pg/pg"

//Connect this function creates and returns a connection to the DB
func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "root",
		Database: "example",
	})
	return db
}
