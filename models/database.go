package models

import (
	"database/sql"
	"os"
	"strings"
)

func InitDB(database, schemaPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}

	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(string(schema))

	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, err
	}

	return db, nil
}
