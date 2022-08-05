package api

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/eleanorhealth/go-exercise/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockUserStore struct {
	domain.UserStorer

	findMock     func(context.Context) ([]*domain.User, error)
	findByIDMock func(context.Context, string) (*domain.User, error)
	saveMock     func(context.Context, *domain.User) error
}

func (m *mockUserStore) Find(ctx context.Context) ([]*domain.User, error) {
	return m.findMock(ctx)
}

func (m *mockUserStore) FindByID(ctx context.Context, id string) (*domain.User, error) {
	return m.findByIDMock(ctx, id)
}

func (m *mockUserStore) Save(ctx context.Context, user *domain.User) error {
	return m.saveMock(ctx, user)
}

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest("", "/", nil)

	entities := []*domain.User{
		{
			ID:        uuid.New().String(),
			NameFirst: "John",
			NameLast:  "Smith",
			Email:     "jsmith@gmail.com",
		},
		{
			ID:        uuid.New().String(),
			NameFirst: "Jane",
			NameLast:  "Doe",
			Email:     "jane.doe@gmail.com",
		},
	}

	expectedRes := &getUsersResponse{
		Users: []*user{
			{
				ID:        entities[0].ID,
				NameFirst: entities[0].NameFirst,
				NameLast:  entities[0].NameLast,
				Email:     entities[0].Email,
			},
			{
				ID:        entities[1].ID,
				NameFirst: entities[1].NameFirst,
				NameLast:  entities[1].NameLast,
				Email:     entities[1].Email,
			},
		},
	}

	userStore := &mockUserStore{}
	userStore.findMock = func(ctx context.Context) ([]*domain.User, error) {
		return entities, nil
	}

	GetUsers(userStore)(rr, r)

	res := &getUsersResponse{}
	err := json.Unmarshal(rr.Body.Bytes(), res)
	assert.NoError(err)

	assert.Equal(expectedRes, res)
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest("", "/", nil)

	entity := &domain.User{
		ID:        uuid.New().String(),
		NameFirst: "John",
		NameLast:  "Smith",
		Email:     "jsmith@gmail.com",
	}

	expectedRes := &user{
		ID:        entity.ID,
		NameFirst: entity.NameFirst,
		NameLast:  entity.NameLast,
		Email:     entity.Email,
	}

	userStore := &mockUserStore{}
	userStore.findByIDMock = func(ctx context.Context, id string) (*domain.User, error) {
		return entity, nil
	}

	GetUserByID(userStore)(rr, r)

	res := &user{}
	err := json.Unmarshal(rr.Body.Bytes(), res)
	assert.NoError(err)

	assert.Equal(expectedRes, res)
}
