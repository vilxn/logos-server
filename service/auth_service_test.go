package service

import (
	"dot/models"
	"testing"
)

func TestRegister(t *testing.T) {
	db, err := models.InitDB("../main.db", "../schema.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepository := models.NewUserRepository(db)
	authService := NewAuthService(*userRepository)

	user := models.User{
		Email:        "helo@gmail.com",
		PasswordHash: "password",
		Role:         models.RoleParent,
	}
	registeredUser, err := authService.Register(user)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Email:", registeredUser.Email)
	t.Log("Password:", registeredUser.PasswordHash)
}
