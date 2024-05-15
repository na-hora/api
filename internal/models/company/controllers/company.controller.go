package companyControllers

import (
	"encoding/json"
	"na-hora/api/internal/dto"
	"na-hora/api/internal/models/company/repositories"
	companyUsecase "na-hora/api/internal/models/company/usecases"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var companyPayload dto.CompanyCreate

	err := json.NewDecoder(r.Body).Decode(&companyPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	appErr := companyPayload.ValidateCNPJ()
	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}

	companyUsecase := companyUsecase.NewCompanyUsecase(repositories.NewCompanyRepository(nil))

	appErr = companyUsecase.CreateCompany(companyPayload)
	if err != nil {
		w.WriteHeader(appErr.StatusCode)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}
}
