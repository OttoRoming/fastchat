package main

import (
	_ "embed"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

var (
	db *sqlx.DB
)

func loadDB() error {
	// avoid shadowing of db
	var err error
	db, err = sqlx.Connect("sqlite3", "database.sqlite3")
	if err != nil {
		return err
	}

	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
}
