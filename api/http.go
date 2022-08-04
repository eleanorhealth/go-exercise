package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondError(w http.ResponseWriter, statusCode int, err error) {
	log.Printf("error: %s\n", err.Error())

	http.Error(w, http.StatusText(statusCode), statusCode)
}

func respond(w http.ResponseWriter, statusCode int, v any) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Printf("error encoding body: %s\n", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
	}
}
