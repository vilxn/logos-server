package handlers

import (
	"dot/models"

	"github.com/gin-gonic/gin"
)

type ParentHandler struct {
	repo *models.ParentRepository
}

func NewParentHandler(repo *models.ParentRepository) *ParentHandler {
	return &ParentHandler{repo: repo}
}

func (h *ParentHandler) GetChildren(c *gin.Context) {
	u, exists := c.Get("userID")
	if !exists {
		c.JSON(400, gin.H{
			"error": "User ID required",
		})
		return
	}

	userID := u.(int64)

	h.repo.SetParentID(userID)
	children, err := h.repo.GetChildren()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, children)
}
