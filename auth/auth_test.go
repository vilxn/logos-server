package auth

import (
	"testing"
)

func TestCreateToken(test *testing.T) {
	userName := "hdsfaafa"
	token, err := CreateToken(userName)

	if err != nil {
		test.Error(err)
	}

	c, err := ValidateToken(token)
	userNameFromToken := c.Username

	if userNameFromToken != userName {
		test.Error("Username token ")
	}
}
