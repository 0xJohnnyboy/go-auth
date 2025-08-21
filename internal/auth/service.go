package auth

import (
	"errors"

	. "goauth/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(username, password string) (*User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, err := HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := User{
		Username: username,
		Password: hashedPassword,
	}

	return &user, s.db.Create(&user).Error
}

func (s *AuthService) Login(username, password string) (*User, error) {
	var user User

	err := s.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	if !VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid password")

	}

	return &user, nil
}
