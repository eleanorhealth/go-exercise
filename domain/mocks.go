package domain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockStore struct {
	mock.Mock
}

func NewMockStore(t *testing.T) *mockStore {
	mock := &mockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (m *mockStore) Users() UserStorer {
	args := m.Called()

	return args.Get(0).(UserStorer)
}

type mockUserStore struct {
	mock.Mock
}

func NewMockUserStore(t *testing.T) *mockUserStore {
	mock := &mockUserStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (m *mockUserStore) Find(ctx context.Context) ([]*User, error) {
	args := m.Called(ctx)

	return args.Get(0).([]*User), args.Error(1)
}

func (m *mockUserStore) FindByID(ctx context.Context, id string) (*User, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*User), args.Error(1)
}

func (m *mockUserStore) Save(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)

	return args.Error(0)
}
