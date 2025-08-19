package auth

import (
	"errors"
	// "fmt"

	. "goauth/internal/models"
	d "goauth/internal/storage"
)

func Register(username, password string) (*User, error) {
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

	db, err := d.Connect()

	if err != nil {
		return nil, err
	}

	return &user, db.Create(&user).Error
}

func Login(username, password string) (*User, error) {
	db, err := d.Connect()

	if err != nil {
		return nil, err
	}

	var user User

	err = db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	if !VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

func Logout() error {
	return errors.New("not implemented")
}
