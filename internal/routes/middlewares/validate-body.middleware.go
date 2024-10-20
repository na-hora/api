package middlewares

import (
	"context"
	"encoding/json"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateStructBody(newStruct interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := newStruct

			if err := json.NewDecoder(r.Body).Decode(target); err != nil {
				utils.ResponseJSON(w, http.StatusBadRequest, "Invalid request body - "+err.Error())
				return
			}

			validate := validator.New()
			if err := validate.Struct(target); err != nil {
				utils.ResponseValidationErrors(err, w, "body")
				return
			}

			ctx := context.WithValue(r.Context(), utils.ValidatedBodyKey, target)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
