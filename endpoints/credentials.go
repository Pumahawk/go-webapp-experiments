package endpoints

import (
	"fmt"
	"net/http"

	"pumahawk.com/webserver/server"
)

func GetCredentialsEndpoint(ctx *server.AppContext) server.EndpointResult {
	db := ctx.DB
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := db.Conn(r.Context())
		if err != nil {
			server.ErroResponse(ctx, w, "Unable to start connections %v", err)
			return
		}
		defer conn.Close()

		rows, err := db.Query("select id from credential limit 1")
		if err != nil {
			server.ErroResponse(ctx, w, "Unable to execute select credential sql query %v", err)
			return
		}
		defer rows.Close()

		var id int64
		if hasOne := rows.Next(); hasOne {
			rows.Scan(&id)
			fmt.Fprintf(w, "Credential id: %d", id)
		} else {
			http.NotFound(w, r)
		}
	}
}

