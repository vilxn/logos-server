package service

import (
	"dot/models"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo models.UserRepository
}

func NewAuthService(userRepo models.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(user models.User) (models.User, error) {
	emailExist, _ := s.userRepo.DoesEmailExist(user.Email)
	if emailExist {
		return models.User{}, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.PasswordHash = string(hash)
	user.CreatedAt = time.Now()

	userID, err := s.userRepo.InsertUser(user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = userID
	return user, nil
}
