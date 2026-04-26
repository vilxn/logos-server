package handlers

import (
	"dot/models"

	"github.com/gin-gonic/gin"
)

type ChildrenHandler struct{
	userRep *models.UserRepository
}

func NewChildrenHandler(ur *models.UserRepository) *ChildrenHandler {
	return &ChildrenHandler{
		userRep: ur,
	}
}

func (h *ChildrenHandler) Get(c *gin.Context) {

}
