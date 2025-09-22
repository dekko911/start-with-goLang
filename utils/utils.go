package utils

// utils use for reuseable func/method, and etc; Over and over. <- another word, this say helper.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]any{
		"status": status,
		"error":  err.Error(),
	})
}

func GetTokenFromRequest(r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization")

	tokenAuth := strings.TrimPrefix(tokenHeader, "Bearer ")
	tokenAuth = strings.TrimSpace(tokenAuth)

	tokenQuery := r.URL.Query().Get("token")

	switch {
	case tokenAuth != "":
		return tokenAuth
	case tokenQuery != "":
		return tokenQuery
	}

	return ""
}
