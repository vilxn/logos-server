package handlers

import (
	"dot/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRep *models.UserRepository
}

func NewUserHandler(ur *models.UserRepository) *UserHandler {
	return &UserHandler{
		userRep: ur,
	}
}

func (uh *UserHandler) GetMe(c *gin.Context) {
	userIdAny, exist := c.Get("userID")
	if !exist {
		c.JSON(400, gin.H{
			"message": "no userID",
		})
		return
	}

	userId, ok := userIdAny.(int64)
	if !ok {
		c.JSON(400, gin.H{
			"message": "couldn't convert to int64",
		})
		return
	}

	user, err := uh.userRep.GetUserByID(userId)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "could not find user with id",
		})
		return
	}

	c.JSON(200, user)
}
