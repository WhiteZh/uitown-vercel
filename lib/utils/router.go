package utils

import "net/http"

type MethodRouter struct {
	Get, Post, Patch, Delete, Put func(w http.ResponseWriter, r *http.Request)
}

func (m *MethodRouter) Route(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if m.Get == nil {
			WriteNotImplementedResponse(w)
			break
		}
		m.Get(w, r)
	case http.MethodPost:
		if m.Post == nil {
			WriteNotImplementedResponse(w)
			break
		}
		m.Post(w, r)
	case http.MethodPatch:
		if m.Patch == nil {
			WriteNotImplementedResponse(w)
			break
		}
		m.Patch(w, r)
	case http.MethodDelete:
		if m.Delete == nil {
			WriteNotImplementedResponse(w)
			break
		}
		m.Delete(w, r)
	case http.MethodPut:
		if m.Put == nil {
			WriteNotImplementedResponse(w)
			break
		}
		m.Put(w, r)
	default:
		WriteNotImplementedResponse(w)
	}
}
