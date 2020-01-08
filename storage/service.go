package storage

import "context"

// Op defines comparison operators
type Op int

const (
	// Eq ==
	Eq Op = iota
	// Gt >
	Gt
	// Ge >=
	Ge
	// Lt <
	Lt
	// Le <=
	Le
	// Ne !=
	Ne
	// In array contains
	In
)

// UpdateAction defines update actions
type UpdateAction int

const (
	// Set sets value
	Set UpdateAction = iota
	// Incr increments value
	Incr
	// Add adds value in array
	Add
	// Remove removes value from array
	Remove
)

const (
	// TypeMissing representing nil entity
	TypeMissing = "MISSING"
)

// Query defines query
type Query struct {
	Field string      `json:"field"`
	Op    Op          `json:"op"`
	Value interface{} `json:"value"`
}

// SortDir query sort direction
type SortDir int

const (
	// Asc -
	Asc SortDir = iota
	// Desc -
	Desc
)

// Limit query limit
type Limit struct {
	From  interface{}
	Limit int
}

// Update -
type Update struct {
	Field  string
	Action UpdateAction
	Value  interface{}
}

// Paged represents paged query result
type Paged struct {
	List     []*E
	HasMore  bool
	NextFrom *E
}

// PagedIDs represents paged ids query result
type PagedIDs struct {
	List     []string
	HasMore  bool
	NextFrom string
}

// Counts -
type Counts map[string]map[string]int64

// Storage the storage interface
type Storage interface {
	DefineEType(ctx context.Context, etype *EType, creationRTypes []*RType) error
	DefineRType(ctx context.Context, rtype *RType) error

	Create(ctx context.Context, e *E, related []*E) (*E, error)
	CreateRelation(ctx context.Context, f, t *E, verb string) error
	RemoveRelation(ctx context.Context, f, t *E, verb string) error
	GetRelation(ctx context.Context, f *E, relation string, limit Limit) (PagedIDs, error)
	HasRelations(ctx context.Context, f, t *E, relations []string) (map[string]bool, error)

	GetByIds(ctx context.Context, typ string, ids []string) ([]*E, error)
	GetByQuery(ctx context.Context, typ string, queries []Query, sorts map[string]SortDir, limit Limit) (Paged, error)
	GetIDsByQuery(ctx context.Context, typ string, queries []Query, sorts map[string]SortDir, limit Limit) (PagedIDs, error)
	UpdateContent(ctx context.Context, typ string, id string, updates []Update) error
	SetState(ctx context.Context, typ string, ids []string, state State) error
	Delete(ctx context.Context, typ string, ids []string) error
	SetMeta(ctx context.Context, typ string, id string, meta map[string]string) error
	DeleteMeta(ctx context.Context, typ string, id string, keys []string) error
	AddTags(ctx context.Context, typ string, id string, tags []string) error
	DeleteTags(ctx context.Context, typ string, id string, tags []string) error

	GetCounts(ctx context.Context, typ string, id string) (Counts, error)
}

// Cache represents cache
type Cache interface {
	Add(es []*E) error
	Remove(es []*E) error
}

// NilOrMissing determines given entity is nil or missing
func NilOrMissing(e *E) bool {
	return e == nil || e.Type == "" || e.Type == TypeMissing
}

// Missing creates an entity representing missing object
func Missing() *E {
	return &E{Type: TypeMissing}
}
