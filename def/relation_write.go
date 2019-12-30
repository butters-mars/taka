package def

// RC creates a RelationCreate instance
func RC() *RelationCreate {
	return &RelationCreate{}
}

// From -
func (rc *RelationCreate) From(f *EX) *RelationCreate {
	rc.From_ = f
	return rc
}

// To -
func (rc *RelationCreate) To(t *EX) *RelationCreate {
	rc.To_ = t
	return rc
}

// Verb -
func (rc *RelationCreate) Verb(v string) *RelationCreate {
	rc.Verb_ = v
	return rc
}

// IsAdd -
func (rc *RelationCreate) IsAdd(v bool) *RelationCreate {
	rc.IsAdd_ = v
	return rc
}

// Can -
func (rc *RelationCreate) Can(user *EX) *RelationCreate {
	rc.User = user
	return rc
}
