package main

import (
	"crypto/rand"
	_ "embed"
	"encoding/base64"

	"github.com/google/uuid"
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

	// db.MustExec(schema)

	return nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func addAccount(username string, password string) error {
	id := uuid.NewString()
	token := uuid.NewString()

	_, err := db.Exec(`INSERT INTO account (id, token, username, password) VALUES (?, ?, ?, ?);`, id, token, username, password)

	return err
}
