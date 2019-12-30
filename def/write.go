package def

// C creates an EntityCreate instance
func C() *EntityCreate {
	return &EntityCreate{}
}

// Col -
func (w *EntityCreate) Col(typ string) *EntityCreate {
	w.Type = typ
	return w
}

// E -
func (w *EntityCreate) E(e *E) *EntityCreate {
	w.E_ = e
	return w
}

// Related -
func (w *EntityCreate) Related(related ...*EX) *EntityCreate {
	w.Related_ = related
	return w
}

// Resources -
func (w *EntityCreate) Resources(resources ...*E) *EntityCreate {
	w.Resources_ = resources
	return w
}

// Can -
func (w *EntityCreate) Can(user *EX, action string) *EntityCreate {
	w.User = user
	w.Action = action
	return w
}
