package def

// Q creates an entity query instance
func U() *EntityUpdate {
	return &EntityUpdate{
		Updates: make([]*Update, 0),
	}
}

// Col defines type to query
func (u *EntityUpdate) Col(typ string) *EntityUpdate {
	u.Type = typ
	return u
}

// Of -
func (u *EntityUpdate) Of(id string) *EntityUpdate {
	u.OfID = id
	return u
}

// Can -
func (u *EntityUpdate) Can(user *EX, action string) *EntityUpdate {
	u.User = user
	u.Action = action
	return u
}

// Set -
func (u *EntityUpdate) Set(field string, value interface{}, typ ValueType) *EntityUpdate {
	u.Updates = append(u.Updates, &Update{
		Field:     field,
		Action:    UpdateAction_Set,
		Value:     toString(value, typ),
		ValueType: typ,
	})

	return u
}

// Incr -
func (u *EntityUpdate) Incr(field string, value int64) *EntityUpdate {
	u.Updates = append(u.Updates, &Update{
		Field:     field,
		Action:    UpdateAction_Incr,
		Value:     toString(value, ValueType_Int64),
		ValueType: ValueType_Int64,
	})

	return u
}

// Add -
func (u *EntityUpdate) Add(field string, value interface{}, typ ValueType) *EntityUpdate {
	u.Updates = append(u.Updates, &Update{
		Field:     field,
		Action:    UpdateAction_Add,
		Value:     toString(value, typ),
		ValueType: typ,
	})

	return u
}

// Remove -
func (u *EntityUpdate) Remove(field string, value interface{}, typ ValueType) *EntityUpdate {
	u.Updates = append(u.Updates, &Update{
		Field:     field,
		Action:    UpdateAction_Remove,
		Value:     toString(value, typ),
		ValueType: typ,
	})

	return u
}

// Updates -
func (u *EntityUpdate) WithUpdates(updates []*Update) *EntityUpdate {
	u.Updates = append(u.Updates, updates...)
	return u
}
