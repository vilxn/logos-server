package auth

import (
	"dot/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mykey")

type UserClaims struct {
	ID   int64           `json:"id"`
	Role models.UserRole `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(user models.User) (string, error) {
	userClaims := UserClaims{
		user.ID,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{},
		func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})

	if err != nil {
		return UserClaims{}, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok {
		return *claims, nil
	} else {
		fmt.Println("Invalid claims")
		return UserClaims{}, fmt.Errorf("Invalid claims")
	}
}
