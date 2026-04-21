package main

import (
	"database/sql"
	"dot/auth"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var username string = "hello"
	tokenString, err := auth.CreateToken(username)
	userClaims, err := auth.ValidateToken(tokenString)

	fmt.Println("Username: ", userClaims.Username)
	fmt.Println("Token: ", tokenString)
}
