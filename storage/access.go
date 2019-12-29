package storage

import "context"

// Allow is a function to check whether a user can do action on entities
type Allow func(ctx context.Context, user *E, action string, entities ...*E) bool

// AccessControl interface of access control for storage
type AccessControl interface {
	Allow(ctx context.Context, user *E, action string, entities ...*E) (bool, error)
}
