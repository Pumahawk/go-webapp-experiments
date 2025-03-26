package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"simpl-go/users-roles/controllers"
	"simpl-go/users-roles/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Start users-roles")
	conf, err := LoadFlags()
	if err != nil {
		log.Fatalf("Invalid flags. %s", err)
	}

	log.Printf("Address: %s", conf.SEAddress)

	err = LoadHttpHandlers(conf)
	if err != nil {
		log.Fatalf("Unable to load http handlers. %s", err)
	}

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server startup error: %w", err)
	}
}

func LoadFlags() (*Config, error) {
	var config Config
	flag.StringVar(&config.SEAddress, "address", ":8080", "Http server address")
	flag.BoolVar(&config.RSIndent, "response-indent", false, "JSON response indentation")
	dbConnS := flag.String("db-conn", "", "Database connection string")
	flag.Parse()
	if *dbConnS == "" {
		return nil, fmt.Errorf("Invalid flags. Database connection string mandatory")
	} else {
		config.DBConnS = *dbConnS
	}
	return &config, nil
}

func LoadHttpHandlers(config *Config) error {
	db, err := GetDB(config)
	if err != nil {
		return fmt.Errorf("Unable load httphandlers for database connection. %w", err)
	}
	jsonHandler := jsonHandlerIndent(config.RSIndent)
	http.HandleFunc("/roles/{id}", BaseChain(db, jsonHandler(controllers.RoleFindById)))
	http.HandleFunc("/identity-attributes/search", BaseChain(db, jsonHandler(controllers.IdentityAttributeSearch)))
	return nil
}

func jsonHandlerIndent(indent bool) func(server.RestController) http.HandlerFunc {
	return func(rc server.RestController) http.HandlerFunc {
		return server.JsonHandler(indent, rc)
	}
}

func DBConnChain(db *sql.DB, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := db.Conn(r.Context())
		if err != nil {
			server.JsonHandler(false, func(r *http.Request) server.RestResponse {
				return controllers.ErrorResponse(500, "Unable to start db connections. %s", err)
			})(w, r)
			return
		}
		defer conn.Close()
		rctx := context.WithValue(r.Context(), server.DBConnK, conn)
		handler(w, r.WithContext(rctx))
	}
}

func LogReqestChain(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("url: %s", r.URL.Path)
		handler(w, r)
	}
}

func BaseChain(db *sql.DB, handler http.HandlerFunc) http.HandlerFunc {
	chain := handler
	chain = DBConnChain(db, chain)
	chain = LogReqestChain(chain)
	return chain
}

type Config struct {
	DBConnS string
	RSIndent bool
	SEAddress string
}

func GetDB(config *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBConnS)
	if err != nil {
		return nil, fmt.Errorf("GetDB error. %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("GetDB Ping error. %w", err)
	}
	return db, nil
}
