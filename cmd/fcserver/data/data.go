package data

import (
	"crypto/rand"
	"database/sql"
	_ "embed"
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

type Account struct {
	ID       string
	Token    string
	Username string
	Password string
}

var (
	db *sqlx.DB
)

func generateToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func LoadDB() error {
	// avoid shadowing of db
	var err error
	db, err = sqlx.Connect("sqlite3", "database.sqlite3")
	if err != nil {
		return err
	}

	_, _ = db.Exec(schema)

	return nil
}

func Close() {
	_ = db.Close()
}

func AddAccount(username string, password string) (Account, error) {
	var account Account
	account.ID = uuid.NewString()

	token, err := generateToken()
	if err != nil {
		return Account{}, err
	}
	account.Token = token

	account.Username = username
	account.Password = password

	_, err = db.Exec(`INSERT INTO account (id, token, username, password) VALUES (?, ?, ?, ?);`, account.ID, account.Token, account.Username, account.Password)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

func GetAccountByUsername(username string) (*Account, error) {
	var account Account

	err := db.Get(&account, `SELECT * FROM account WHERE username = ?`, username)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &account, err
}
