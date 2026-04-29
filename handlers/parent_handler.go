package handlers

import (
	"dot/models"

	"github.com/gin-gonic/gin"
)

type AddChildRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	Notes     string `json:"notes"`
}

type ParentHandler struct {
	repo *models.ParentRepository
}

func NewParentHandler(repo *models.ParentRepository) *ParentHandler {
	return &ParentHandler{repo: repo}
}

func (h *ParentHandler) GetChildren(c *gin.Context) {
	u, _ := c.Get("userID")
	parentID := u.(int64)

	children, err := h.repo.GetChildren(parentID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, children)
}

func (h *ParentHandler) AddChild(c *gin.Context) {
	u, _ := c.Get("userID")
	parentID := u.(int64)

	var req AddChildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	childID, err := h.repo.AddChild(parentID, models.Child{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: req.BirthDate,
		Notes:     req.Notes,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"childID": childID,
	})
}
