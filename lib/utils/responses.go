package utils

import (
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func SetContentTypeJSON(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) {

	SetContentTypeJSON(w)
	w.WriteHeader(statusCode)

	EncodeJSONOrPanic(w, ErrorResponse{
		ErrorMessage: message,
	})
}

//func WriteInternalErrorResponse(w http.ResponseWriter) {
//	if err := WriteErrorResponse(w, "Internal Server Error", http.StatusInternalServerError); err != nil {
//		log.Fatal(err)
//	}
//}

func WriteBadRequestResponse(w http.ResponseWriter) {
	WriteErrorResponse(w, "Bad Request", http.StatusBadRequest)
}

func WriteNotImplementedResponse(w http.ResponseWriter) {
	WriteErrorResponse(w, "Not Implemented", http.StatusNotImplemented)
}

func WriteUnauthorizedResponse(w http.ResponseWriter) {
	WriteErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
}
