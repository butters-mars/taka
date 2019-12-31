package storage

import "fmt"

// Obj provides common methods for entity & relation
type Obj interface {
	GetId() int64
	GetType() string
	GetBiztype() string
	GetState() int32
	GetCTime() int64
	GetUTime() int64
}

var (
	// ErrUnauthenticated -
	ErrUnauthenticated = fmt.Errorf("Unauthenticated")
	// ErrUnauthorized -
	ErrUnauthorized = fmt.Errorf("Unauthorized")
	// ErrUnimplemented  -
	ErrUnimplemented = fmt.Errorf("Unimplemented")
)
