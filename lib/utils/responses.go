package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func SetContentTypeJSON(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) error {

	SetContentTypeJSON(w)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(ErrorResponse{
		ErrorMessage: message,
	})
}

//func WriteInternalErrorResponse(w http.ResponseWriter) {
//	if err := WriteErrorResponse(w, "Internal Server Error", http.StatusInternalServerError); err != nil {
//		log.Fatal(err)
//	}
//}

func WriteBadRequestResponse(w http.ResponseWriter) {
	if err := WriteErrorResponse(w, "Bad Request", http.StatusBadRequest); err != nil {
		log.Fatal(err)
	}
}

func WriteNotImplementedResponse(w http.ResponseWriter) {
	if err := WriteErrorResponse(w, "Not Implemented", http.StatusNotImplemented); err != nil {
		log.Fatal(err)
	}
}
