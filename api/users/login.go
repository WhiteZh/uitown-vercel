package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"uitown-vercel/lib/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var db = utils.TryConnectDB(w)
	defer utils.TryCloseDB(db, w)

	var queries = r.URL.Query()

	var email = queries.Get("email")
	var passwordHashed = queries.Get("password_hashed")

	if email == "" || passwordHashed == "" {
		utils.WriteBadRequestResponse(w)
		return
	}

	var id int
	var associatedPasswordHashed string

	var row = db.QueryRow(`SELECT id, password_hashed FROM users WHERE email = $1`, email)

	if err := row.Scan(&id, &associatedPasswordHashed); err != nil {
		utils.WriteInternalErrorResponse(w)
		log.Fatal(err)
	}

	var res int

	if associatedPasswordHashed != passwordHashed {
		res = -1
	} else {
		res = id
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		utils.WriteInternalErrorResponse(w)
		log.Fatal(err)
	}
}
