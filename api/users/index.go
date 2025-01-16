package users

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"maps"
	"net/http"
	"strings"
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

		row := db.QueryRow(`SELECT name, email, password_hashed, aboutme, icon, icon_type FROM users WHERE id = $1`, queryId)

		qres := struct {
			Name           string
			Email          string
			PasswordHashed string
			Aboutme        string
			Icon           *[]byte
			IconType       *string
		}{}
		{
			err := row.Scan(&qres.Name, &qres.Email, &qres.PasswordHashed, &qres.Aboutme, &qres.Icon, &qres.IconType)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					utils.WriteErrorResponse(w, "User ID does not exist", http.StatusNotFound)
					return
				}
				log.Panic(err)
			}
		}

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
			IconType       *string `json:"icon_type"`
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
			icon := base64.StdEncoding.EncodeToString(*qres.Icon)
			res.Icon = &icon
			res.IconType = qres.IconType
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrPanic(w, res)
	},
	Patch: func(w http.ResponseWriter, r *http.Request) {

		params := struct {
			Id                int     `json:"id"`
			PasswordHashed    string  `json:"password_hashed"`
			Name              *string `json:"name"`
			Email             *string `json:"email"`
			NewPasswordHashed *string `json:"new_password_hashed"`
			Aboutme           *string `json:"aboutme"`
			Icon              *string `json:"icon"`
			IconType          *string `json:"icon_type"`
		}{}
		{
			err := json.NewDecoder(r.Body).Decode(&params)
			if err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}
		}
		{
			cnt := 0
			if params.Icon == nil {
				cnt++
			}
			if params.IconType == nil {
				cnt++
			}
			if cnt == 1 {
				utils.WriteBadRequestResponse(w)
				return
			}
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		{
			row := utils.QueryRowDBOrPanic(db, `SELECT password_hashed FROM users WHERE id = $1`, params.Id)
			var passwordHashed string
			utils.ScanOrPanic(row, &passwordHashed)

			if passwordHashed != params.PasswordHashed {
				utils.WriteUnauthorizedResponse(w)
				return
			}
		}

		modColumns := make(map[string]any)
		if params.Name != nil {
			modColumns["name"] = params.Name
		}
		if params.Email != nil {
			modColumns["email"] = params.Email
		}
		if params.NewPasswordHashed != nil {
			modColumns["password_hashed"] = params.NewPasswordHashed
		}
		if params.Aboutme != nil {
			modColumns["aboutme"] = params.Aboutme
		}
		if params.Icon != nil {
			if params.IconType == nil {
				log.Panicln("miracle happened")
			}
			modColumns["icon"] = params.Icon
			modColumns["icon_type"] = params.IconType
		}

		var updateClause string
		updateValues := make([]any, len(modColumns))
		{
			updateSegments := make([]string, len(modColumns))
			i := 0
			for v := range maps.Keys(modColumns) {
				updateSegments[i] = fmt.Sprintf("%s = $%d", v, i+1)
				updateValues[i] = modColumns[v]
				i++
			}
			updateClause = strings.Join(updateSegments, ", ")
		}

		{
			_, err := db.Exec(fmt.Sprintf("UPDATE users SET %s WHERE id = %d", updateClause, params.Id), updateValues...)
			if err != nil {
				utils.WriteErrorResponse(
					w,
					"Unprocessable Entity; most likely due to the new name or email conflict with some existing one(s)",
					http.StatusUnprocessableEntity)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	},
}
