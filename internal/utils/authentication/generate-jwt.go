package authentication

import (
	"na-hora/api/internal/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateToken(ID uuid.UUID, username string) (string, *utils.AppError) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      ID,
		"username": username,
		"iss":      "Na Hora",
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":      time.Now().Unix(),
	})

	jwtSecret := viper.Get("JWT_SECRET").(string)
	token, err := claims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", &utils.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return token, nil
}
