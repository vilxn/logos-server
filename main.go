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
	parentRepository := models.NewParentRepository(db)

	authService := service.NewAuthService(*userRepository)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userRepository)
	parentHandler := handlers.NewParentHandler(parentRepository)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Get rid of in release
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
	parent := api.Group("/parent")
	{
		parent.Use(middleware.RoleMiddleware(models.RoleParent))
		parent.GET("/children", parentHandler.GetChildren)
		parent.POST("/child", parentHandler.AddChild)
	}
	r.Run(":8080")
}
