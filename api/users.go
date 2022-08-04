package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/eleanorhealth/go-exercise/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type createUserRequest struct {
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Email     string `json:"email"`
}

func CreateUser(userStore domain.UserStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := &createUserRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			respondError(w, http.StatusBadRequest, fmt.Errorf("decoding request body: %w", err))
			return
		}

		user := &domain.User{
			ID:        uuid.New().String(),
			NameFirst: req.NameFirst,
			NameLast:  req.NameLast,
			Email:     req.Email,
		}

		err = userStore.Save(ctx, user)
		if err != nil {
			respondError(w, http.StatusBadRequest, fmt.Errorf("saving user: %w", err))
			return
		}

		respond(w, http.StatusOK, nil)
	}
}

type user struct {
	ID        string `json:"id"`
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Email     string `json:"email"`
}

type getUsersResponse struct {
	Users []*user `json:"users"`
}

func GetUsers(userStore domain.UserStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		users, err := userStore.Find(ctx)
		if err != nil {
			respondError(w, http.StatusInternalServerError, fmt.Errorf("finding users: %w", err))
			return
		}

		res := &getUsersResponse{}

		for _, entity := range users {
			res.Users = append(res.Users, &user{
				ID:        entity.ID,
				NameFirst: entity.NameFirst,
				NameLast:  entity.NameLast,
				Email:     entity.Email,
			})
		}

		respond(w, http.StatusOK, res)
	}
}

func GetUserByID(userStore domain.UserStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")

		entity, err := userStore.FindByID(ctx, id)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				respondError(w, http.StatusNotFound, err)
				return
			}

			respondError(w, http.StatusInternalServerError, fmt.Errorf("finding users: %w", err))
			return
		}

		user := &user{
			ID:        entity.ID,
			NameFirst: entity.NameFirst,
			NameLast:  entity.NameLast,
			Email:     entity.Email,
		}

		respond(w, http.StatusOK, user)
	}
}
