package main

import (
	"dot/models"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := models.InitDB("main.db", "schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := models.NewUserRepository(db)

	user := models.User{
		Email:        "test2@gmail.com",
		PasswordHash: "test",
		FirstName:    "test",
		LastName:     "test",
		Role:         models.RoleParent,
		CreatedAt:    time.Now(),
	}

	err = userRepository.InsertUser(user)
	if err != nil {
		panic(err)
	}
}
