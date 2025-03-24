package endpoints

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"pumahawk.com/webserver/server"
)

type KeyPairDB struct {
	Id string
	PublicKey []byte
	PrivateKey []byte
}

type KeyPairInfo struct {
	Id string `json:"id"`
}

func DownloadPublicKeyEndpoint(actx *server.AppContext) server.EndpointResult {
	db := actx.DB
	log := actx.Log
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err :=  db.Conn(r.Context())
		if err != nil {
			log.Error("Unable to connect to database: %v", err)
			w.WriteHeader(500)
			return
		}
		kp, err := GetKeyPairDB(r.Context(), conn)
		if err != nil {
			log.Error("DB KeyPair error: %s", err)
			w.WriteHeader(500)
			return
		}
		w.Write(kp.PublicKey)
	}
}

func GetInfoPrivateKeyEndpoint(actx *server.AppContext) server.EndpointResult {
	db := actx.DB
	return server.JsonEndpoint(func(r *http.Request) server.HttpResponse {
		conn, err :=  db.Conn(r.Context())
		if err != nil {
			return DBRestError(err)
		}
		kp, err := GetKeyPairDB(r.Context(), conn)
		if err != nil {
			return RestError(500, "DB KeyPair error: %s", err)
		}
		kpinfo := KeyPairInfo{
			Id: kp.Id,
		}
		return server.HttpResponse{
			Code: 200,
			Body: kpinfo,
		}
	})
}

func DownloadPrivateKeyEndpoint(actx *server.AppContext) server.EndpointResult {
	db := actx.DB
	log := actx.Log
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err :=  db.Conn(r.Context())
		if err != nil {
			log.Error("Unable to connect to database: %v", err)
			w.WriteHeader(500)
			return
		}
		kp, err := GetKeyPairDB(r.Context(), conn)
		if err != nil {
			log.Error("DB KeyPair error: %s", err)
			w.WriteHeader(500)
			return
		}
		w.Write(kp.PrivateKey)
	}
}

func GetKeyPairDB(ctx context.Context, conn *sql.Conn) (*KeyPairDB, error) {
	sql := `
	select 
		id,
		public_key,
		private_key
	from keypair
	limit 1
	`
	row := conn.QueryRowContext(ctx, sql)
	var kp KeyPairDB
	err := row.Scan(&kp.Id, &kp.PublicKey, &kp.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve keypair: %w", err)
	}
	return &kp, nil
}
