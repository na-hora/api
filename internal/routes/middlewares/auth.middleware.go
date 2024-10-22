package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type AuthService interface {
	JwtAuthMiddleware(next http.Handler) http.Handler
	GetClaimsFromContext(ctx context.Context) jwt.MapClaims
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

type contextKey string

const (
	TokenContextKey contextKey = "userClaims"
)

func (m *authService) JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		_, claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		if claims["iss"] != "Na Hora" {
			http.Error(w, fmt.Sprintf("invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), TokenContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	jwtSecret := viper.Get("JWT_SECRET").(string)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, fmt.Errorf("invalid token")
}

func (m *authService) GetClaimsFromContext(ctx context.Context) jwt.MapClaims {
	if claims, ok := ctx.Value(TokenContextKey).(jwt.MapClaims); ok {
		return claims
	}
	return nil
}
