package storage

import "context"

// Model defines model to be stored
type Model interface {
	GetId() int64
	GetCt() int64
	GetUt() int64
	GetState() int32
}

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

// Sort -
type Sort struct {
	Field string
	Dir   SortDir
}

// Limit query limit
type Limit struct {
	From  interface{}
	Limit int
	Page  int
}

// Update -
type Update struct {
	Field  string
	Action UpdateAction
	Value  interface{}
}

// Service defines methods of a storage
type Service interface {
	Create(ctx context.Context, typ string, m Model) error
	Update(ctx context.Context, typ string, q Query, updates []Update) error
	GetByID(ctx context.Context, typ string, id int64, out interface{}) error
	GetByIDs(ctx context.Context, typ string, ids []int64, out interface{}) error
	GetOneByQuery(ctx context.Context, typ string, qs []Query, obj interface{}) error
	GetByQuery(ctx context.Context, typ string, qs []Query, sorts []Sort, limit Limit, slice interface{}) (bool, error)
	SetState(ctx context.Context, typ string, id int64, state int32) error

	GetCount(ctx context.Context, typ string, id int64) (map[string]int64, error)
	IncrCount(ctx context.Context, typ string, id int64, delta map[string]int64) error
	AddRelation(ctx context.Context, typ string, id, to int64, rel string) error
	RemoveRelation(ctx context.Context, typ string, id, to int64, rel string) error
	HasRelation(ctx context.Context, typ string, id, to int64, rel string) (bool, error)

	GetRelated(ctx context.Context, typ string, id int64, rel string, limit Limit) ([]int64, error)

	// UpdateByID(id string, update interface{}) error
	// GetByIDs(ids []int64) ([]Model, error)
	// GetByQuery()
	// Delete(m Model) error

	// CreateRelation(m Model, to string, relation string, allowMultiple bool) error
	// HasRelation(m, Model, to string, relation string) (bool, error)
	//GetRelated(f Model, relation string, reversed bool)
}

/*
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
*/

/*
    rpc create(CreateEWithRsReq) returns (E);
    rpc createRelation(RelationReq) returns (Empty);
    rpc removeRelation(RelationReq) returns (Empty);
    rpc hasRelations(HasRelationsReq) returns (HasRelations);
    rpc getRelation(GetRelationReq) returns (PagedIDs);

    rpc getByIds(GetByIDsReq) returns (EList);
    rpc getByQuery(GetByQueryReq) returns (Paged);
    rpc updateContent(UpdateContentReq) returns (Empty);
    rpc setState(SetStateReq) returns (Empty);
    rpc delete(DeleteReq) returns (Empty);
    rpc setMeta(SetMetaReq) returns (Empty);
	rpc deleteMeta(DeleteMetaReq) returns (Empty);
	rpc addTags(UpdateTagsReq) returns (Empty);
	rpc deleteTags(UpdateTagsReq) returns (Empty);

	rpc getCounts(GetCountsReq) returns (Counts);
*/
