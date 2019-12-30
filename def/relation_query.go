package def

// RQ creates a relationQuery instance
func RQ() *RelationQuery {
	return &RelationQuery{
		PopOp: &PopulationOption{},
	}
}

// Col defines type to query
func (rq *RelationQuery) Col(typ string) *RelationQuery {
	rq.Type = typ
	return rq
}

// Rel -
func (rq *RelationQuery) Rel(relation string) *RelationQuery {
	rq.Relation = relation
	return rq
}

// Of -
func (rq *RelationQuery) Of(of *EX) *RelationQuery {
	rq.OfID = of
	return rq
}

// PopCounts populates counts
func (rq *RelationQuery) PopCounts() *RelationQuery {
	rq.PopOp.WithCounts = true
	return rq
}

// PopResources -
func (rq *RelationQuery) PopResources() *RelationQuery {
	rq.PopOp.WithResources = true
	return rq
}

// PopRelated -
func (rq *RelationQuery) PopRelated(idx int, typ string) *RelationQuery {
	rq.PopOp.Related[int32(idx)] = typ
	return rq
}

// PopRelationWith -
func (rq *RelationQuery) PopRelationWith(e *EX) *RelationQuery {
	if e != nil {
		rq.PopOp.HasRelationsWith = e.Entity

	}
	return rq
}

// PopRelationOf -
func (rq *RelationQuery) PopRelationOf(rels ...string) *RelationQuery {
	rq.PopOp.HasRelationsOf = append(rq.PopOp.HasRelationsOf, rels...)
	return rq
}

// From -
func (rq *RelationQuery) From(from string) *RelationQuery {
	rq.FromID = from
	return rq
}

// Limit -
func (rq *RelationQuery) Limit(limit int32) *RelationQuery {
	rq.WithLimit = limit
	return rq
}
