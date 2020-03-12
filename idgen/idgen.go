package idgen

type Service interface {
	GenID() (int64, error)
}
