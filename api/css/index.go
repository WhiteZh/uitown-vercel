package css

import (
	"encoding/json"
	"fmt"
	"net/http"
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
			row := utils.QueryRowDBOrPanic(db, "SELECT password_hashed FROM users WHERE id = $1", params.UserID)

			var realPasswordHashed string
			utils.ScanOrPanic(row, &realPasswordHashed)

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
}
