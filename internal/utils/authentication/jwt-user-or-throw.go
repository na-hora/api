package authentication

import (
	"context"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	authentication "na-hora/api/internal/routes/middlewares"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
)

type UserLogged struct {
	ID        uuid.UUID
	Username  string
	CompanyID uuid.UUID
}

func JwtUserOrThrow(ctx context.Context) (*UserLogged, *utils.AppError) {
	authService := authentication.NewAuthService()
	claims := authService.GetClaimsFromContext(ctx)
	ID := claims["sub"]

	if ID == nil {
		return nil, &utils.AppError{
			Message:    "Invalid token",
			StatusCode: http.StatusUnauthorized,
		}
	}

	parsedID, parseErr := uuid.Parse(ID.(string))
	if parseErr != nil {
		return nil, &utils.AppError{
			Message:    "Invalid token",
			StatusCode: http.StatusUnauthorized,
		}
	}

	userService := injector.InitializeUserService(config.DB)

	userFound, err := userService.GetByID(parsedID)

	if err != nil {
		return nil, err
	}

	return &UserLogged{
		ID:        userFound.ID,
		Username:  userFound.Username,
		CompanyID: userFound.CompanyID,
	}, nil
}
