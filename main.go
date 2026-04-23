package main

import (
	"dot/handlers"
	"dot/models"
	"dot/service"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := models.InitDB("main.db", "schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := models.NewUserRepository(db)

	authService := service.NewAuthService(*userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()

	api := r.Group("/logos")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
