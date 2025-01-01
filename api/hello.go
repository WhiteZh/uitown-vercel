package api

import (
	"fmt"
	"net/http"
	"uitown-vercel/lib/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s from %s", utils.HelloWorld, r.URL.RawQuery)
}
