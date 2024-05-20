package handlers

import (
	"encoding/json"
	"fmt"
	tokenDTOs "na-hora/api/internal/models/token/dtos"
	"na-hora/api/internal/models/token/services"
	"na-hora/api/internal/utils"
	"net/http"
	"os"

	"github.com/go-playground/validator"
)

type TokenHandler interface {
	GenerateRegisterLink(w http.ResponseWriter, r *http.Request)
}

type tokenHandler struct {
	tokenService services.TokenService
}

func GetTokenHandler() TokenHandler {
	tokenService := services.GetTokenService()
	return &tokenHandler{
		tokenService,
	}
}

func (th *tokenHandler) GenerateRegisterLink(w http.ResponseWriter, r *http.Request) {
	var tokenPayload tokenDTOs.GenerateTokenRequestBody

	err := json.NewDecoder(r.Body).Decode(&tokenPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, "Invalid body")
		return
	}

	validate := validator.New()
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
		URL: fmt.Sprintf("%s/register/%s", os.Getenv("API_PUBLIC_URL"), token.Key),
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
