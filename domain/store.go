package domain

import "context"

type Storer interface {
	Users() UserStorer
}

type UserStorer interface {
	Find(context.Context) ([]*User, error)
	FindByID(context.Context, string) (*User, error)
	Save(context.Context, *User) error
}
