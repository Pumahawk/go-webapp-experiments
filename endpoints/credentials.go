package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"pumahawk.com/webserver/server"
)

type Credential struct {
	Id            int64
	Content       []byte
	ParticipantId *string
}

func GetCredentialsEndpoint(ctx *server.AppContext) server.EndpointResult {
	db := ctx.DB
	log := ctx.Log
	tpl := ctx.Template
	return func(w http.ResponseWriter, r *http.Request) {
		errorResponse := server.ErrorResponseFunc(ctx, w)
		conn, err := db.Conn(r.Context())
		if err != nil {
			errorResponse("Unable to start connections %v", err)
			return
		}
		defer conn.Close()

		ids := r.PathValue("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			errorResponse("Invalid id parameter %v", err)
			return
		}
		credentials, err := CredentialsMetadataDBFindById(ctx, id)
		if err != nil {
			log.Error("Unabe to retieve error", err)
			errorResponse("Unable to retrieve credentials from database %v", id)
			return
		}

		if credentials != nil {
			tpl.ExecuteTemplate(w, "credential.tmpl.html", credentials)
		} else {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Credential not found credential %d", id)
		}
	}
}

func CredentialsMetadataDBFindById(ctx *server.AppContext, id int) (*Credential, error) {
	log := ctx.Log
	db := ctx.DB
	log.Debug("Find credential by id %s", id)
	row := db.QueryRow(`
		select
			id,
			participant_id
		from credential
		where
			id = $1
		`, id)
	var credential Credential
	err := row.Scan(&credential.Id, &credential.ParticipantId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, fmt.Errorf("Unable to find credential by id %w", err)
		}
	}
	return &credential, nil
}
