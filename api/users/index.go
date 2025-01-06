package users

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"uitown-vercel/lib/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methodRouter.Route(w, r)
}

var methodRouter = utils.MethodRouter{
	Get: func(w http.ResponseWriter, r *http.Request) {

		var queryId int
		var queryPasswordHashed string

		queries := r.URL.Query()
		err := utils.UnescapeQueryValues(queries)
		if err != nil {
			utils.WriteBadRequestResponse(w)
			return
		}

		{
			rawQueryID := queries.Get("id")
			if rawQueryID == "" {
				utils.WriteBadRequestResponse(w)
				return
			}

			_, err := fmt.Sscanf(rawQueryID, "%d", &queryId)
			if err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}

			queryPasswordHashed = queries.Get("password_hashed")
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		row := utils.QueryRowDBOrPanic(db, `SELECT name, email, password_hashed, aboutme, icon, icon_type FROM users WHERE id = $1`, queryId)

		qres := struct {
			Name           string
			Email          string
			PasswordHashed string
			Aboutme        string
			Icon           *[]byte
			IconType       *string
		}{}
		utils.ScanOrPanic(row, &qres.Name, &qres.Email, &qres.PasswordHashed, &qres.Aboutme, &qres.Icon, &qres.IconType)

		if qres.PasswordHashed != queryPasswordHashed {
			utils.WriteUnauthorizedResponse(w)
			return
		}

		type Response struct {
			Id             int     `json:"id"`
			Name           string  `json:"name"`
			Email          string  `json:"email"`
			PasswordHashed string  `json:"password_hashed"`
			Aboutme        string  `json:"aboutme"`
			Icon           *string `json:"icon"`
		}

		res := Response{
			Id:             queryId,
			Name:           qres.Name,
			Email:          qres.Email,
			PasswordHashed: queryPasswordHashed,
			Aboutme:        qres.Aboutme,
		}

		if qres.Icon != nil {
			if qres.IconType == nil {
				log.Panic("users.icon is not NULL while users.icon_type is")
			}
			dataURL := fmt.Sprintf("data:%s;base64,%s", qres.IconType, base64.StdEncoding.EncodeToString(*qres.Icon))
			res.Icon = &dataURL
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrPanic(w, res)
	},
}
