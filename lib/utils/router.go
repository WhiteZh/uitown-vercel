package utils

import (
	"fmt"
	"log"
	"net/http"
)

type MethodRouter struct {
	Get, Post, Patch, Delete, Put HandlerType
	MiddleWares                   []HandlerType
}

func notImplementedHandler(w http.ResponseWriter, _ *http.Request) {
	WriteNotImplementedResponse(w)
}

func (m *MethodRouter) Route(w http.ResponseWriter, r *http.Request) {

	log.Println(fmt.Sprintf("%s %s", r.Method, r.URL.RequestURI()))

	if m.MiddleWares != nil {
		for _, middleWare := range m.MiddleWares {
			middleWare(w, r)
		}
	}

	handler := notImplementedHandler

	switch r.Method {
	case http.MethodGet:
		handler = m.Get
	case http.MethodPost:
		handler = m.Post
	case http.MethodPatch:
		handler = m.Patch
	case http.MethodDelete:
		handler = m.Delete
	case http.MethodPut:
		handler = m.Put
	}

	handler(w, r)
}
