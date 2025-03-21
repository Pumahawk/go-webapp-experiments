package main

import (
	"database/sql"
	"log"
	"net/http"

	"pumahawk.com/webserver/database"
	"pumahawk.com/webserver/endpoints"
	mylog "pumahawk.com/webserver/log"
	"pumahawk.com/webserver/server"
)

func main() {
	err := LoadAppFlags()
	if err != nil {
		log.Fatal("Unable to load global flags", err)
	}

	InitApp()

	log.Printf("Start web server %s", GlobalAppFlag.Address)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Unable to start server", err)
	}
}

func InitApp() {
	context := CreateAppContext()
	InitEndpoints(&context);
}

func BaseChain(endpointFunc server.EndpointResult) server.EndpointResult {
	var chain server.EndpointResult = endpointFunc
	chain = LogInterceptor(endpointFunc)
	return chain
}

func LogInterceptor(endpointFunc server.EndpointResult) server.EndpointResult {
	return func(r http.ResponseWriter, rq *http.Request) {
		path := rq.URL.Path
		log.Printf("request: %s", path)
		endpointFunc(r, rq)
	}
}

func CreateAppContext() server.AppContext {
	logger := CreateLogger()
	db := CreateDB()
	ctx := server.AppContext{
		Log: &logger,
		DB: db,
	}
	return ctx
}

func InitEndpoints(ctx *server.AppContext) {
	helloWolrdEndpoint := endpoints.HelloWorlsEndpoint(ctx)

	http.HandleFunc("/", BaseChain(helloWolrdEndpoint))
}

func CreateLogger() mylog.Logger {
	logger := mylog.Logger{}
	return logger
}

func CreateDB() *sql.DB {
	dbConf := GetDatabaseConfiguration()
	db, err := database.CreateDatabaseConnection(dbConf)
	if err != nil {
		log.Fatal("Unable to start database connection", err)
	}
	return db
}

func GetDatabaseConfiguration() database.DBConf {
	return database.DBConf{
		User: GlobalAppFlag.DB.User,
		Password: GlobalAppFlag.DB.Password,
		DBName: GlobalAppFlag.DB.Database,
	}
}
