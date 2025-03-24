package endpoints

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"pumahawk.com/webserver/server"
)

type Credential struct {
	Id            int64   `json:"id"`
	Content       []byte  `json:"content"`
	ParticipantId *string `json:"participantId"`
}

type CredentialMetadata struct {
	Id            int64   `json:"id"`
	ParticipantId *string `json:"participantId"`
}

type GetCredentialParameter struct {
	Id int
}

func ExtractGetCredentialParameter(r *http.Request) (*GetCredentialParameter, error) {
	ids := r.PathValue("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return nil, err
	}
	return &GetCredentialParameter{
		Id: id,
	}, nil
}

func GetCredentialsEndpoint(actx *server.AppContext) server.EndpointResult {
	log := actx.Log
	return server.JsonEndpoint(func(r *http.Request) server.HttpResponse {
		params, err := ExtractGetCredentialParameter(r)
		if err != nil {
			return RestError(500, "Invalid id parameter %v", err)
		}

		conn, err := actx.DB.Conn(r.Context())
		if err != nil {
			return DBRestError(err)
		}

		credentials, err := CredentialsMetadataDBFindById(r.Context(), actx, conn, params.Id)
		if err != nil {
			log.Error("Unabe to retieve error", err)
			return RestError(500, "Unable to retrieve credentials from database %v", params.Id)
		}

		return server.HttpResponse{
			Code: 200,
			Body: credentials,
		}
	})
}

func DonwloadCredentialEndpoint(actx *server.AppContext) server.EndpointResult {
	log := actx.Log
	return func(w http.ResponseWriter, r *http.Request) {
		ids := r.PathValue("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			log.Error("Unable to extract id. %v", err)
			w.WriteHeader(400)
			return
		}
		content, err := CredentialsContentDBFindById(actx, r.Context(), id)
		if err != nil {
			log.Error("Unable to retrieve credentials content, %v", err)
			w.WriteHeader(500)
			return
		} else if content != nil {
			w.WriteHeader(200)
			w.Write(content)
			return
		} else {
			w.WriteHeader(404)
			return
		}
	}
}

func CredentialsMetadataDBFindById(ctx context.Context, actx *server.AppContext, conn *sql.Conn, id int) (*CredentialMetadata, error) {
	log := actx.Log
	log.Debug("Find credential by id %s", id)
	row := conn.QueryRowContext(ctx, `
		select
			id,
			participant_id
		from credential
		where
			id = $1
		`, id)
	var credential CredentialMetadata
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

func CredentialsContentDBFindById(actx *server.AppContext, ctx context.Context, id int) ([]byte, error) {
	log := actx.Log
	db := actx.DB
	log.Debug("Find credential content by id %s", id)
	row := db.QueryRowContext(ctx, `
		select
			content
		from credential
		where
			id = $1
		`, id)
	var content []byte
	err := row.Scan(&content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, fmt.Errorf("Unable to find credential content by id %w", err)
		}
	}
	return content, nil
}
