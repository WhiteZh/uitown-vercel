package utils

import "net/http"

type HandlerType func(w http.ResponseWriter, r *http.Request)
