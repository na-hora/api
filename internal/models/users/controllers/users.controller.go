package userControllers

import (
	"encoding/json"
	"net/http"

	dto "na-hora/api/internal/dto"
	authentication "na-hora/api/internal/utils/authentication"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user dto.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*
	   buscar o usuário no banco de dados
	*/

	token, appError := authentication.CreateToken(user.Username)
	if err != nil {
		w.WriteHeader(appError.StatusCode)
		json.NewEncoder(w).Encode(appError.Message)
		return
	}

	response := dto.LoginResponse{
		Token:    token,
		Username: user.Username,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
