package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS users ( id serial PRIMARY KEY NOT NULL, phone TEXT NOT NULL, password TEXT NOT NULL);")
	fmt.Println("ASdsadasd")
	if err != nil {
		return err
	}
	return nil
}

func Down00004(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
