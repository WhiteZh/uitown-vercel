package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) error {

	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(ErrorResponse{
		ErrorMessage: message,
	})
}

func WriteInternalErrorResponse(w http.ResponseWriter) {
	if err := WriteErrorResponse(w, "Internal Server Error", http.StatusInternalServerError); err != nil {
		log.Fatal(err)
	}
}

func WriteBadRequestResponse(w http.ResponseWriter) {
	if err := WriteErrorResponse(w, "Bad Request", http.StatusBadRequest); err != nil {
		log.Fatal(err)
	}
}
