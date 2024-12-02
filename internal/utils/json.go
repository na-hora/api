package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if v == "EOF" {
		v = "invalid input"
	}

	return json.NewEncoder(w).Encode(v)
}
