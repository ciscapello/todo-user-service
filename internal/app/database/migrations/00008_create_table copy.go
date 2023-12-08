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
ALTER TABLE users ADD COLUMN created_at TIMESTAMP DEFAULT NOW();
ALTER TABLE users ADD COLUMN updated_at TIMESTAMP DEFAULT NOW();
`

func Up(tx *sql.Tx) error {
	_, err := tx.Exec(queryString)
	// _, err := tx.Exec("CREATE TABLE IF NOT EXISTS users ( id serial PRIMARY KEY NOT NULL, phone TEXT NOT NULL, password TEXT NOT NULL);")
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
