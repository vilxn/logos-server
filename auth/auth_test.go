package auth

import (
	"dot/models"
	"testing"
	"time"
)

func TestCreateToken(test *testing.T) {
	user := models.User{
		ID:        12,
		Email:     "test@gmail.com",
		Role:      models.RoleParent,
		CreatedAt: time.Now(),
	}
	token, err := CreateToken(user)
	if err != nil {
		test.Error(err)
	}

	c, err := ValidateToken(token)
	if err != nil {
		test.Error(err)
	}

	test.Log("ID from token: ", c.ID)
	test.Log("Role from token:  ", c.Role)
}
