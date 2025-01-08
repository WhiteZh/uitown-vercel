package users

import (
	"encoding/base64"
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
		queries := r.URL.Query()
		{
			err := utils.UnescapeQueryValues(queries)
			if err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}
		}

		params := struct {
			id                int
			passwordHashed    string
			name              *string
			email             *string
			newPasswordHashed *string
			description       *string
			icon              *string
			iconType          *string
		}{}

		{
			_, err := fmt.Sscanf(queries.Get("id"), "%d", &params.id)
			if err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}
		}
		{
			params.passwordHashed = queries.Get("password_hashed")
			if params.passwordHashed == "" {
				utils.WriteBadRequestResponse(w)
				return
			}
		}
		{
			name := queries.Get("name")
			if name != "" {
				params.name = &name
			}
		}
		{
			email := queries.Get("email")
			if email != "" {
				params.email = &email
			}
		}
		{
			newPasswordHashed := queries.Get("new_password_hashed")
			if newPasswordHashed != "" {
				params.newPasswordHashed = &newPasswordHashed
			}
		}
		{
			description := queries.Get("description")
			if description != "" {
				params.description = &description
			}
		}
		{
			icon := queries.Get("icon")
			iconType := queries.Get("icon_type")
			if icon != "" {
				if iconType == "" {
					utils.WriteBadRequestResponse(w)
					return
				}
				params.icon = &icon
				params.iconType = &iconType
			}
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		{
			row := utils.QueryRowDBOrPanic(db, `SELECT password_hashed FROM users WHERE id = $1`, params.id)
			var passwordHashed string
			utils.ScanOrPanic(row, &passwordHashed)

			if passwordHashed != params.passwordHashed {
				utils.WriteUnauthorizedResponse(w)
				return
			}
		}

		modColumns := make(map[string]any)
		if params.name != nil {
			modColumns["name"] = params.name
		}
		if params.email != nil {
			modColumns["email"] = params.email
		}
		if params.newPasswordHashed != nil {
			modColumns["password_hashed"] = params.newPasswordHashed
		}
		if params.description != nil {
			modColumns["password_hashed"] = params.description
		}
		if params.icon != nil {
			if params.iconType == nil {
				log.Panicln("miracle happened")
			}
			modColumns["icon"] = params.icon
			modColumns["icon_type"] = params.iconType
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
			_, err := db.Exec(fmt.Sprintf("UPDATE users SET %s WHERE id = %d", updateClause, params.id), updateValues...)
			if err != nil {
				utils.WriteErrorResponse(
					w,
					"Unprocessable Entity; most likely because the new name or email conflicts with some existing one(s)",
					http.StatusUnprocessableEntity)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	},
}
