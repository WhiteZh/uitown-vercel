package valid

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"uitown-vercel/lib/types"
	"uitown-vercel/lib/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methodRouter.Route(w, r)
}

var methodRouter = utils.MethodRouter{
	Get: func(w http.ResponseWriter, r *http.Request) {

		type orderType int
		const (
			orderById orderType = iota
			orderByName
			orderByViewedTime
		)

		params := struct {
			category *types.CssCategoryType
			authorId *int
			limit    *int
			offset   *int
			order    *[]orderType
		}{}

		{
			queries := r.URL.Query()

			{
				rawCategory := queries.Get("category")
				if rawCategory != "" {
					category, err := types.ConvertStringToCssCategory(rawCategory)
					if err != nil {
						utils.WriteBadRequestResponse(w)
						return
					}
					params.category = &category
				}
			}
			{
				rawAuthorId := queries.Get("author_id")
				if rawAuthorId != "" {
					var authorId int
					_, err := fmt.Sscanf(rawAuthorId, "%d", &authorId)
					if err != nil {
						utils.WriteBadRequestResponse(w)
						return
					}
					params.authorId = &authorId
				}
			}
			{
				rawLimit := queries.Get("limit")
				if rawLimit != "" {
					var limit int
					_, err := fmt.Sscanf(rawLimit, "%d", &limit)
					if err != nil || limit < 1 {
						utils.WriteBadRequestResponse(w)
						return
					}
					params.limit = &limit
				}
			}
			{
				rawOffset := queries.Get("offset")
				if rawOffset != "" {
					var offset int
					_, err := fmt.Sscanf(rawOffset, "%d", &offset)
					if err != nil || offset < 0 {
						utils.WriteBadRequestResponse(w)
						return
					}
					params.offset = &offset
				}
			}
			{
				rawOrders := queries["order"]
				if len(rawOrders) != 0 {
					orders := make([]orderType, len(rawOrders))
					for i, v := range rawOrders {
						switch v {
						case "id":
							orders[i] = orderById
						case "name":
							orders[i] = orderByName
						case "viewed_time":
							orders[i] = orderByViewedTime
						default:
							utils.WriteBadRequestResponse(w)
							return
						}
					}
					params.order = &orders
				}
			}
		}

		var sqlClause = "SELECT id FROM css"

		{
			chains := make([]string, 0, 2)

			if params.category != nil {
				chains = append(chains, fmt.Sprintf("category = '%s'", types.ConvertCssCategoryToString(*params.category)))
			}

			if params.authorId != nil {
				chains = append(chains, fmt.Sprintf("author_id = %d", *params.authorId))
			}

			if len(chains) > 0 {
				whereClause := fmt.Sprintf(" WHERE %s", strings.Join(chains, " AND "))
				sqlClause += whereClause
			}
		}

		if params.order != nil {
			order := make([]string, len(*params.order))
			for i, v := range *params.order {
				switch v {
				case orderById:
					order[i] = "id"
				case orderByName:
					order[i] = "name"
				case orderByViewedTime:
					order[i] = "viewed_time"
				default:
					log.Panic("uncovered `orderType`")
				}
			}

			orderByClause := fmt.Sprintf(" ORDER BY %s", strings.Join(order, ", "))
			sqlClause += orderByClause
		}

		if params.limit != nil {
			limitClause := fmt.Sprintf(" LIMIT %d", *params.limit)
			sqlClause += limitClause
		}

		if params.offset != nil {
			offsetClause := fmt.Sprintf(" OFFSET %d", *params.offset)
			sqlClause += offsetClause
		}

		db := utils.ConnectDBOrPanic()
		defer utils.CloseDBOrPanic(db)

		res := make([]int, 0)

		rows := utils.QueryDBOrPanic(db, sqlClause)
		defer utils.CloseRowsOrPanic(rows)
		for rows.Next() {
			var id int
			utils.ScanOrPanic(rows, &id)
			res = append(res, id)
		}

		utils.SetContentTypeJSON(w)
		utils.EncodeJSONOrPanic(w, res)
	},
}
