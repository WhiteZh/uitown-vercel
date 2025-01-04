package css

import (
	"fmt"
	"net/http"
	"uitown-vercel/lib/utils"

	"github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methodRouter.Route(w, r)
}

var methodRouter = utils.MethodRouter{
	Get: func(w http.ResponseWriter, r *http.Request) {

		var queryIds = r.URL.Query()["id"]
		var ids = make([]int, len(queryIds))
		for i, v := range queryIds {
			if _, err := fmt.Sscanf(v, "%d", &ids[i]); err != nil {
				utils.WriteBadRequestResponse(w)
				return
			}
		}

		var db = utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		var rows = utils.QueryDBOrPanic(db, `SELECT css.id, css.name, css.viewed_time, css.author_id, css.html, css.css, css.category FROM css WHERE id = ANY($1)`, pq.Array(ids))

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
}
