package main

import (
	"log"
	"net/http"

	"github.com/eleanorhealth/go-exercise/api"
	"github.com/eleanorhealth/go-exercise/storage/memory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	httpAddr = ":3000"
)

func main() {
	userStore := memory.NewUserStore()

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/users", api.CreateUser(userStore))
		r.Get("/users", api.GetUsers(userStore))
		r.Get("/users/{id}", api.GetUserByID(userStore))
	})

	log.Printf("Starting HTTP server at %s\n", httpAddr)
	err := http.ListenAndServe(httpAddr, r)
	if err != nil {
		log.Fatal(err)
	}
}
