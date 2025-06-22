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

func BadRequest(w http.ResponseWriter, err error) {
	RespondJSON(w, http.StatusBadRequest, map[string]string{
		"error": err.Error(),
	})
}

func ParseIDParam(r *http.Request, param string) (int, error) {
	return strconv.Atoi(chi.URLParam(r, param))
}

func RespondMessage(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, models.NewMessageResponse(message))
}
