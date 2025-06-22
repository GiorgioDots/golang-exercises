package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/giorgiodots/todo-go-api/models"
	"github.com/go-chi/chi/v5"
)

func RespondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func ExtractJSON[T any](r *http.Request) (T, error) {
	var zero T

	if r.Body == nil {
		return zero, errors.New("empty request body")
	}
	defer r.Body.Close()

	var value T
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&value); err != nil {
		return zero, err
	}

	if decoder.More() {
		return zero, errors.New("unexpected data after JSON object")
	}

	return value, nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondError(w http.ResponseWriter, code int, err error) {
	RespondJSON(w, code, ErrorResponse{Error: err.Error()})
}

func ParseIDParam(r *http.Request, param string) (int, error) {
	p := chi.URLParam(r, param)
	return strconv.Atoi(p)
}

func RespondMessage(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, models.NewMessageResponse(message))
}
