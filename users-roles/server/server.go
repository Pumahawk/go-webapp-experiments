package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type RestController = func(r *http.Request) RestResponse

type RestResponse struct {
	Code int
	Body any
}

func JsonHandler(indent bool, rest RestController) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := rest(r)
		select {
		case _ = <- r.Context().Done():
			log.Println("Context closed. Unable to write response.")
		default:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(resp.Code)
			encoder := json.NewEncoder(w)
			if indent {
				encoder.SetIndent("", "    ")
			}
			err := encoder.Encode(resp)
			if err != nil {
				log.Printf("Unable to write json response: %v", err)
			}
		}
	}
}
