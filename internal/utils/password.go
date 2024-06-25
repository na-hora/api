package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordUtilInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type passwordUtil struct{}

func GetPasswordUtil() PasswordUtilInterface {
	return &passwordUtil{}
}

func (pu *passwordUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("error hashing password")
	}
	return string(bytes), nil
}

func (pu *passwordUtil) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
