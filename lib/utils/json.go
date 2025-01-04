package utils

import (
	"encoding/json"
	"io"
	"log"
)

func EncodeJSONOrPanic(w io.Writer, a any) {
	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		log.Panic(err)
	}
}
