package storage

// Obj provides common methods for entity & relation
type Obj interface {
	GetId() int64
	GetType() string
	GetBiztype() string
	GetState() int32
	GetCTime() int64
	GetUTime() int64
}
