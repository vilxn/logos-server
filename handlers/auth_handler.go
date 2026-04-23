package handlers

import (
	"dot/auth"
	"dot/models"
	"dot/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required,min=6"`
	Role     models.UserRole `json:"role" binding:"required"`

	FirstName string `json:"first-name" binding:"required"`
	LastName  string `json:"last-name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		fmt.Print(req)
		return
	}

	user, err := h.service.Register(models.User{
		Email:        req.Email,
		PasswordHash: req.Password,
		Role:         req.Role,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
	})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.CreateToken(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(201, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"token": tokenString,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(models.User{Email: req.Email, PasswordHash: req.Password})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}
