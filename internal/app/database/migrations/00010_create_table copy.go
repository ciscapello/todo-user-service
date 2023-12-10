package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up, Down)
}

var queryString = `
CREATE TABLE IF NOT EXISTS users ( 
	id serial PRIMARY KEY NOT NULL, 
	username TEXT NOT NULL, 
	password TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW());

CREATE TABLE IF NOT EXISTS sessions (
	id serial PRIMARY KEY NOT NULL,
	refresh_token text,
	expired_at 
)
`

func Up(tx *sql.Tx) error {
	_, err := tx.Exec(queryString)
	fmt.Println("ASdsadasd")
	if err != nil {
		return err
	}
	return nil
}

func Down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
