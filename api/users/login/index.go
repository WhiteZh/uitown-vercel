package login

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"uitown-vercel/lib/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methodRouter.Route(w, r)
}

var methodRouter = utils.MethodRouter{
	Get: func(w http.ResponseWriter, r *http.Request) {

		var queries = r.URL.Query()
		err := utils.UnescapeQueryValues(queries)
		if err != nil {
			utils.WriteBadRequestResponse(w)
			return
		}

		var email = queries.Get("email")
		var passwordHashed = queries.Get("password_hashed")

		if email == "" || passwordHashed == "" {
			utils.WriteBadRequestResponse(w)
			return
		}

		var db = utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		var id int
		var associatedPasswordHashed string

		row := db.QueryRow(`SELECT id, password_hashed FROM users WHERE email = $1`, email)

		{
			err := row.Scan(&id, &associatedPasswordHashed)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					utils.SetContentTypeJSON(w)
					utils.EncodeJSONOrPanic(w, -1)
					return
				}
				log.Panic(err)
			}
		}

		var res int

		if associatedPasswordHashed != passwordHashed {
			res = -1
		} else {
			res = id
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrPanic(w, res)
	},
}
