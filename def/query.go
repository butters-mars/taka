package def

import (
	fmt "fmt"
	"strings"
)

var opMap = map[Op]string{
	Op_Eq: "==",
	Op_Lt: "<",
	Op_Gt: ">",
	Op_Le: "<=",
	Op_Ge: ">=",
	Op_Ne: "!=",
	Op_In: "IN",
}

// Q creates an entity query instance
func Q() *EntityQuery {
	return &EntityQuery{
		Qt:    QueryType_Paging,
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
	s := toString(val, typ)
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

func toString(val interface{}, typ ValueType) string {
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

	return s
}

func (q *EntityQuery) ToSQL() string {
	var buf strings.Builder

	buf.WriteString("SELECT FROM ")
	buf.WriteString(q.Type)
	buf.WriteString(" ")

	if q.Qt == QueryType_Paging {
		first := true
		for _, q := range q.Queries {
			if first {
				buf.WriteString("WHERE ")
				first = false
			} else {
				buf.WriteString("AND ")
			}

			key := strings.ToLower(q.Field)
			buf.WriteString(key)
			buf.WriteString(" ")
			buf.WriteString(opMap[q.Op])
			buf.WriteString(" ")
			buf.WriteString(fmt.Sprintf("%v", q.Value))
			buf.WriteString(" ")
		}

		if len(q.Sorts) > 0 {
			buf.WriteString("ORDER BY ")
		}

		for k, dir := range q.Sorts {
			d := "ASC"
			_dir := 1
			if dir == SortDir_Desc {
				d = "DESC"
				_dir = -1
			}

			if first {
				first = false
			} else {
			}

			key := strings.ToLower(k)
			if _dir == -1 {
				key = fmt.Sprintf("-%s", key)
			}

			buf.WriteString(strings.ToLower(k))
			buf.WriteString(" ")
			buf.WriteString(d)
			buf.WriteString(" ")
		}

		buf.WriteString("LIMIT ")
		buf.WriteString(fmt.Sprintf("%d", q.WithLimit))
		buf.WriteString(" ")
	} else if q.Qt == QueryType_One {
		buf.WriteString("WHERE id = ")
		buf.WriteString(q.Id)
	} else if q.Qt == QueryType_ByIds {
		buf.WriteString("WHERE id IN (")
		buf.WriteString(strings.Join(q.Ids, ","))
		buf.WriteString(")")
	}

	return buf.String()
}
