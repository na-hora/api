package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *AppError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", &AppError{
			Message:    "invalid password",
			StatusCode: http.StatusInternalServerError,
		}
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
