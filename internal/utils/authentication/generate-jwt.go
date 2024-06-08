package authentication

import (
	"na-hora/api/internal/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(username string) (string, *utils.AppError) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "Na Hora",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
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
