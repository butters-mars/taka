package storage

// Sequencer -
type Sequencer interface {
	Next(key string) (int64, error)
}
