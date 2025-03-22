package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	"pumahawk.com/webserver/database"
	"pumahawk.com/webserver/endpoints"
	mylog "pumahawk.com/webserver/log"
	"pumahawk.com/webserver/server"
	"pumahawk.com/webserver/templates"
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
	InitEndpoints(&context)
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
	tpl := GetTemplate()
	ctx := server.AppContext{
		Log:      &logger,
		DB:       db,
		Template: tpl,
	}
	return ctx
}

func InitEndpoints(ctx *server.AppContext) {
	helloWolrdEndpoint := endpoints.HelloWorlsEndpoint(ctx)
	getCredentialsEndpoint := endpoints.GetCredentialsEndpoint(ctx)

	http.HandleFunc("/credentials/{id}", BaseChain(getCredentialsEndpoint))
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
	user, _ := os.LookupEnv("DBUSER")
	password, _ := os.LookupEnv("DBPASSWORD")
	dbname, _ := os.LookupEnv("DBDBNAME")
	return database.DBConf{
		User:     user,
		Password: password,
		DBName:   dbname,
	}
}

func GetTemplate() *template.Template {
	return templates.LoadTemplateOrFatal()
}
