package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	if err != nil {
		log.Printf("error serving %s: %s\n", r.URL.Path, err.Error())
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func respondJSON(w http.ResponseWriter, statusCode int, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
