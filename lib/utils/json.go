package utils

import (
	"encoding/json"
	"io"
	"log"
)

func EncodeJSONOrFatal(w io.Writer, a any) {
	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		log.Fatal(err)
	}
}
