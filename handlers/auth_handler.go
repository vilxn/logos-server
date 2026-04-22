package handlers

import (
	"dot/models"
	"dot/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required,min=6"`
	Role     models.UserRole `json:"role" binding:"required"`
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
		return
	}

	user, err := h.service.Register(models.User{Email: req.Email, PasswordHash: req.Password, Role: req.Role})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}
