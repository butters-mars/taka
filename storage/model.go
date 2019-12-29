package storage

import "fmt"

// CountType how a relation increments a counter
type CountType int

const (
	// CountBoth adds counter for both ends
	CountBoth CountType = iota
	// CountFromEnd adds counter form the from end
	CountFromEnd
	// CountToEnd adds counter form the from end
	CountToEnd
	// CountNone does not add counter
	CountNone
)

// State entities' state
type State int

const (
	// StateDeleted deleted
	StateDeleted State = iota
	// StatePrivate owner visible
	StatePrivate
	// State1 -
	State1
	// State2 -
	State2
	// State3 -
	State3
	// StateFriend friend visible
	StateFriend
	// StatePublic everyone visible
	StatePublic
)

// RType defines a relation type
type RType struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Verb      string    `json:"verb"`
	Multiple  bool      `json:"multiple"`
	CountType CountType `json:"count_type" yaml:"count_type"`
	//Simple    bool      `json:"simple"`
}

// EType defines an entity type
type EType struct {
	Name string `json:"name"`
	//RTypes []*RType `json:"rtypes"`
	//SimpleRTypes []*RType `json:"simple_rtypes"`
}

// E represents an entity
type E struct {
	Type      string            `json:"type"`
	ID        string            `json:"id"`
	ID1       string            `json:"id1"`
	ID2       string            `json:"id2"`
	ID3       string            `json:"id3"`
	ID4       string            `json:"id4"`
	CTime     int64             `json:"ctime"`
	UTime     int64             `json:"utime"`
	Score     int64             `json:"score"`
	Score1    int64             `json:"score1"`
	State     State             `json:"state"`
	Tags      []string          `json:"tags"`
	Meta      map[string]string `json:"meta"`
	Resources []string          `json:"resources"`
	Content   map[string]string `json:"content"`
	//Text      string                 `json:"text"`
	//Bytes     []byte                 `json:"bytes"`
}

func (e E) String() string {
	return fmt.Sprintf("E[%v](%v,%v,%v,%v state=%v)", e.Type, e.ID, e.ID1, e.ID2, e.ID3, e.State)
}
