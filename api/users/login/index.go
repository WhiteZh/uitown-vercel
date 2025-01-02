package handler

import (
	"net/http"
	"uitown-vercel/lib/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methodRouter.Route(w, r)
}

var methodRouter = utils.MethodRouter{
	Get: func(w http.ResponseWriter, r *http.Request) {

		var queries = r.URL.Query()

		var email = queries.Get("email")
		var passwordHashed = queries.Get("password_hashed")

		if email == "" || passwordHashed == "" {
			utils.WriteBadRequestResponse(w)
			return
		}

		var db = utils.ConnectDBOrFatal()
		defer utils.CloseDBOrFatal(db)

		var id int
		var associatedPasswordHashed string

		var row = utils.QueryRowDBOrFatal(db, `SELECT id, password_hashed FROM users WHERE email = $1`, email)

		utils.ScanOrFatal(row, &id, &associatedPasswordHashed)

		var res int

		if associatedPasswordHashed != passwordHashed {
			res = -1
		} else {
			res = id
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrFatal(w, res)
	},
}
