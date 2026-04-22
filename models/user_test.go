package models

import (
	"database/sql"
	"testing"
)

var createdUserId int64

var ur *UserRepository

func TestInsertUser(test *testing.T) {
	db, _ := sql.Open("sqlite3", "../main.db")
	ur = NewUserRepository(db)

}
