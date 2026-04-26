package main

import (
	"dot/handlers"
	"dot/middleware"
	"dot/models"
	"dot/service"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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
	userHandler := handlers.NewUserHandler(userRepository)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/logos")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
	user := api.Group("/user")
	{
		user.Use(middleware.AuthMiddleware())
		user.GET("/me", userHandler.GetMe)
	}
	children := api.Group("/children")
	{
		children.Use(middleware.RoleMiddleware(models.RoleParent, models.RoleSpecialist))
		children.GET("")
	}
	r.Run(":8080")
}
