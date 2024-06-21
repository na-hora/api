package handlers

import (
	"encoding/json"
	"fmt"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	tokenDTOs "na-hora/api/internal/models/token/dtos"
	"na-hora/api/internal/models/token/services"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type TokenHandlerInterface interface {
	GenerateRegisterLink(w http.ResponseWriter, r *http.Request)
}

type TokenHandler struct {
	tokenService services.TokenServiceInterface
}

func GetTokenHandler() TokenHandlerInterface {
	tokenService := injector.InitializeTokenService(config.DB)
	return &TokenHandler{
		tokenService,
	}
}

func (th *TokenHandler) GenerateRegisterLink(w http.ResponseWriter, r *http.Request) {
	var tokenPayload tokenDTOs.GenerateTokenRequestBody

	err := json.NewDecoder(r.Body).Decode(&tokenPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, "Invalid body")
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(tokenPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w)
		return
	}

	token, sErr := th.tokenService.Generate(tokenPayload)
	if sErr != nil {
		utils.ResponseJSON(w, sErr.StatusCode, sErr.Message)
		return
	}

	response := &tokenDTOs.GenerateTokenResponse{
		URL: fmt.Sprintf("%s/company/register?validator=%s", viper.Get("WEB_URL"), token.Key),
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
