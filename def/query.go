package def

import (
	fmt "fmt"
)

// Q creates an entity query instance
func Q() *EntityQuery {
	return &EntityQuery{
		PopOp: &PopulationOption{},
	}
}

// Col defines type to query
func (q *EntityQuery) Col(typ string) *EntityQuery {
	q.Type = typ
	return q
}

// PopCounts populates counts
func (q *EntityQuery) PopCounts() *EntityQuery {
	q.PopOp.WithCounts = true
	return q
}

// PopResources -
func (q *EntityQuery) PopResources() *EntityQuery {
	q.PopOp.WithResources = true
	return q
}

// PopRelated -
func (q *EntityQuery) PopRelated(idx int, typ string) *EntityQuery {
	q.PopOp.Related[int32(idx)] = typ
	return q
}

// PopRelationWith -
func (q *EntityQuery) PopRelationWith(e *EX) *EntityQuery {
	if e != nil {
		q.PopOp.HasRelationsWith = e.Entity
	}
	return q
}

// PopRelationOf -
func (q *EntityQuery) PopRelationOf(rels ...string) *EntityQuery {
	q.PopOp.HasRelationsOf = append(q.PopOp.HasRelationsOf, rels...)
	return q
}

// Can -
func (q *EntityQuery) Can(user *EX, action string) *EntityQuery {
	q.User = user
	q.Action = action
	return q
}

// From -
func (q *EntityQuery) From(from string) *EntityQuery {
	q.FromID = from
	return q
}

// Limit -
func (q *EntityQuery) Limit(limit int32) *EntityQuery {
	q.WithLimit = limit
	return q
}

// Where -
func (q *EntityQuery) Where(field string, op Op, val interface{}, typ ValueType) *EntityQuery {
	s := ""
	switch typ {
	case ValueType_Int, ValueType_Int64:
		s = fmt.Sprintf("%d", val)
	case ValueType_Double:
		s = fmt.Sprintf("%f", val)
	case ValueType_Bool:
		s = fmt.Sprintf("%t", val)
	case ValueType_Bytes, ValueType_String:
		s = fmt.Sprintf("%s", val)
	default:
		s = fmt.Sprintf("%v", val)
	}

	qry := &Query{
		Field:     field,
		Op:        op,
		Value:     s,
		ValueType: typ,
	}
	q.Queries = append(q.Queries, qry)
	return q
}

// Sort -
func (q *EntityQuery) Sort(field string, dir SortDir) *EntityQuery {
	q.Sorts[field] = dir
	return q
}

// ID -
func (q *EntityQuery) ID(id string) *EntityQuery {
	q.Id = id
	q.Qt = QueryType_One
	return q
}
