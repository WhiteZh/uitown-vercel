package css

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"uitown-vercel/lib/types"
	"uitown-vercel/lib/utils"

	"github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methodRouter.Route(w, r)
}

var methodRouter = utils.MethodRouter{
	Get: func(w http.ResponseWriter, r *http.Request) {

		queries := r.URL.Query()
		err := utils.UnescapeQueryValues(queries)
		if err != nil {
			utils.WriteBadRequestResponse(w)
			return
		}

		var queryIds = queries["id"]
		var ids = make([]int, len(queryIds))
		for i, v := range queryIds {
			if _, err := fmt.Sscanf(v, "%d", &ids[i]); err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}
		}

		var db = utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		var rows = utils.QueryDBOrPanic(db,
			`SELECT css.id, css.name, css.viewed_time, css.author_id, css.html, css.css, css.category FROM css WHERE id = ANY($1)`,
			pq.Array(ids))
		defer utils.CloseRowsOrPanic(rows)

		type Response struct {
			Id         int    `json:"id"`
			Name       string `json:"name"`
			ViewedTime int    `json:"viewed_time"`
			AuthorId   int    `json:"author_id"`
			HTML       string `json:"html"`
			CSS        string `json:"css"`
			Category   string `json:"category"`
		}

		var res = make([]Response, len(ids))

		for i := 0; rows.Next(); i++ {
			t := Response{}

			utils.ScanOrPanic(rows, &t.Id, &t.Name, &t.ViewedTime, &t.AuthorId, &t.HTML, &t.CSS, &t.Category)

			res[i] = t
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrPanic(w, res)
	},
	Post: func(w http.ResponseWriter, r *http.Request) {

		var params struct {
			UserID         int
			PasswordHashed string
			Name           string
			Html           string
			Css            string
			Category       types.CssCategoryType
		}
		{
			body := struct {
				UserID         int    `json:"userID"`
				PasswordHashed string `json:"password_hashed"`
				Name           string `json:"name"`
				Html           string `json:"html"`
				Css            string `json:"css"`
				Category       string `json:"category"`
			}{}
			{
				err := json.NewDecoder(r.Body).Decode(&body)
				if err != nil {
					utils.WriteBadRequestResponse(w)
					return
				}
			}
			{
				category, err := types.ConvertStringToCssCategory(body.Category)
				if err != nil {
					utils.WriteBadRequestResponse(w)
					return
				}
				params.Category = category
			}
			params.UserID = body.UserID
			params.PasswordHashed = body.PasswordHashed
			params.Name = body.Name
			params.Html = body.Html
			params.Css = body.Css
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		{
			row := db.QueryRow("SELECT password_hashed FROM users WHERE id = $1", params.UserID)

			var realPasswordHashed string
			{
				err := row.Scan(&realPasswordHashed)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						utils.WriteErrorResponse(w, "Unprocessable response; user id does not exist", http.StatusUnprocessableEntity)
						return
					}
					log.Panic(err)
				}
			}

			if realPasswordHashed != params.PasswordHashed {
				utils.WriteUnauthorizedResponse(w)
				return
			}
		}

		var newID int64
		{
			row := utils.QueryRowDBOrPanic(db,
				"INSERT INTO css (name, html, css, category, author_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				params.Name, params.Html, params.Css, params.Category, params.UserID)

			utils.ScanOrPanic(row, &newID)
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrPanic(w, newID)
	},
	Patch: func(w http.ResponseWriter, r *http.Request) {

		var params struct {
			Id             int
			PasswordHashed string
			Name           *string
			Html           *string
			Css            *string
			Category       *types.CssCategoryType
		}
		// initialize `params`
		{
			body := struct {
				Id             int     `json:"id"`
				PasswordHashed string  `json:"password_hashed"`
				Name           *string `json:"name"`
				Html           *string `json:"html"`
				Css            *string `json:"css"`
				Category       *string `json:"category"`
			}{}
			{
				err := json.NewDecoder(r.Body).Decode(&body)
				if err != nil {
					utils.WriteBadRequestResponse(w)
					return
				}
			}

			params.Id = body.Id
			params.PasswordHashed = body.PasswordHashed
			params.Name = body.Name
			params.Html = body.Html
			params.Css = body.Css
			if body.Category != nil {
				category, err := types.ConvertStringToCssCategory(*body.Category)
				if err != nil {
					utils.WriteBadRequestResponse(w)
					return
				}
				params.Category = &category
			}
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		// validate authentication
		{
			var realPasswordHashed string
			// retrieve `realPasswordHashed`
			{
				row := db.QueryRow("SELECT password_hashed FROM users WHERE id = (SELECT author_id FROM css WHERE id = $1)", params.Id)
				log.Println(params.Id)

				err := row.Scan(&realPasswordHashed)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						utils.WriteUnauthorizedResponse(w)
						return
					}
					log.Panic(err)
				}
			}

			if params.PasswordHashed != realPasswordHashed {
				utils.WriteUnauthorizedResponse(w)
				return
			}
		}

		// db update
		{
			var sqlClause string
			var columnValues []any
			// calculate and assign the value of `updateClause` & `updateValues`
			{
				var sqlKeys []string
				cnt := 2 // start with 2 as $1 is id

				if params.Name != nil {
					sqlKeys = append(sqlKeys, fmt.Sprintf("name = $%d", cnt))
					cnt++
					columnValues = append(columnValues, *params.Name)
				}

				if params.Html != nil {
					sqlKeys = append(sqlKeys, fmt.Sprintf("html = $%d", cnt))
					cnt++
					columnValues = append(columnValues, *params.Html)
				}

				if params.Css != nil {
					sqlKeys = append(sqlKeys, fmt.Sprintf("css = $%d", cnt))
					cnt++
					columnValues = append(columnValues, *params.Css)
				}

				if params.Category != nil {
					sqlKeys = append(sqlKeys, fmt.Sprintf("category = $%d", cnt))
					cnt++
					columnValues = append(columnValues, types.ConvertCssCategoryToString(*params.Category))
				}

				sqlClause = fmt.Sprintf("UPDATE css SET %s WHERE id = $1", strings.Join(sqlKeys, ", "))
			}

			// if `columnValues` is never updated/be inserted, it will not be initialized and stay nil (as the default value)
			if columnValues != nil {
				sqlValues := []any{params.Id}
				sqlValues = append(sqlValues, columnValues...)
				utils.ExecDBOrPanic(db, sqlClause, sqlValues...)
			}
		}

		w.WriteHeader(http.StatusOK)
	},
	Delete: func(w http.ResponseWriter, r *http.Request) {

		queries := r.URL.Query()
		{
			err := utils.UnescapeQueryValues(queries)
			if err != nil {
				utils.WriteUnauthorizedResponse(w)
				return
			}
		}

		params := struct {
			Id             int
			PasswordHashed string
		}{
			PasswordHashed: queries.Get("password_hashed"),
		}
		{
			var id int
			_, err := fmt.Sscan(queries.Get("id"), &id)
			if err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}
			params.Id = id
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		{
			row := utils.QueryRowDBOrPanic(db, "SELECT password_hashed FROM users WHERE id = (SELECT author_id FROM css WHERE id = $1)", params.Id)

			var realPasswordHashed string
			err := row.Scan(&realPasswordHashed)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					utils.WriteErrorResponse(w, "Invalid CSS id or its associated author id", http.StatusUnprocessableEntity)
					return
				}
				log.Panic(err)
				return
			}

			if realPasswordHashed != params.PasswordHashed {
				utils.WriteUnauthorizedResponse(w)
				return
			}
		}

		utils.ExecDBOrPanic(db, "DELETE FROM css WHERE id = $1", params.Id)

		w.WriteHeader(http.StatusOK)
	},
}
