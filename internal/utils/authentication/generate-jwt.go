package authentication

import (
	"na-hora/api/internal/utils"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, *utils.AppError) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "Na Hora",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", &utils.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return token, nil
}
