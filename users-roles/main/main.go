package main

import (
	"flag"
	"log"
	"net/http"
	"simpl-go/users-roles/controllers"
	"simpl-go/users-roles/server"
)

func main() {
	log.Println("Start users-roles")
	conf := LoadFlags()
	log.Printf("Address: %s", conf.Address)

	LoadHttpHandlers(conf)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server startup error: %w", err)
	}
}

func LoadFlags() (config Config) {
	flag.StringVar(&config.Address, "address", ":8080", "Http server address")
	flag.BoolVar(&config.Indent, "response-indent", false, "JSON response indentation")
	flag.Parse()
	return
}

func LoadHttpHandlers(config Config) {
	jsonHandler := jsonHandlerIndent(config.Indent)
	http.HandleFunc("/roles/{id}", jsonHandler(controllers.RoleFindById))
}

func jsonHandlerIndent(indent bool) func(server.RestController)func(http.ResponseWriter, *http.Request) {
	return func(rc server.RestController) func(http.ResponseWriter, *http.Request) {
		return server.JsonHandler(indent, rc)
	}
}

type Config struct {
	Address string
	Indent bool
}
