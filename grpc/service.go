package grpc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/butters-mars/taka/def"
	"github.com/butters-mars/taka/storage"
	"github.com/butters-mars/tiki/logging"
)

// Wrapper provides grpc wrapper for storage service
type Wrapper struct {
	storage storage.Storage
}

// NewWrapper -
func NewWrapper(storage storage.Storage) *Wrapper {
	return &Wrapper{
		storage: storage,
	}
}

// DefineEType implements StorageServer method
func (s Wrapper) DefineEType(ctx context.Context, req *def.DefineETypeReq) (*def.Empty, error) {
	rts := make([]*storage.RType, 0)
	for _, rt := range req.GetCreationRTypes() {
		rts = append(rts, &storage.RType{
			From:      rt.From,
			To:        rt.To,
			Verb:      rt.Verb,
			Multiple:  rt.Multiple,
			CountType: storage.CountType(rt.CountType),
		})
	}
	err := s.storage.DefineEType(ctx, &storage.EType{
		Name: req.EType.GetName(),
	}, rts)
	return &def.Empty{}, err
}

// DefineRType implements StorageServer method
func (s Wrapper) DefineRType(ctx context.Context, req *def.RType) (*def.Empty, error) {
	err := s.storage.DefineRType(ctx, &storage.RType{
		From:      req.From,
		To:        req.To,
		Verb:      req.Verb,
		Multiple:  req.Multiple,
		CountType: storage.CountType(req.CountType),
	})
	return &def.Empty{}, err
}

// Create implements StorageServer method
func (s Wrapper) Create(ctx context.Context, req *def.CreateEWithRsReq) (*def.E, error) {
	e := toModel(req.E)

	related := make([]*storage.E, 0)
	for _, e := range req.Related {
		related = append(related, toModel(e))
	}

	e, err := s.storage.Create(ctx, e, related)
	if err != nil {
		return nil, err
	}

	return toDef(e), nil
}

// CreateRelation implements StorageServer method
func (s Wrapper) CreateRelation(ctx context.Context, req *def.RelationReq) (*def.Empty, error) {
	f, e := toModel(req.From), toModel(req.To)
	err := s.storage.CreateRelation(ctx, f, e, req.Verb)
	return &def.Empty{}, err
}

// RemoveRelation implements StorageServer method
func (s Wrapper) RemoveRelation(ctx context.Context, req *def.RelationReq) (*def.Empty, error) {
	f, e := toModel(req.From), toModel(req.To)
	err := s.storage.RemoveRelation(ctx, f, e, req.Verb)
	return &def.Empty{}, err
}

// GetRelation implements StorageServer method
func (s Wrapper) GetRelation(ctx context.Context, req *def.GetRelationReq) (*def.PagedIDs, error) {
	f := toModel(req.From)
	rel := req.Relation
	limit := toLimit(req.Limit)

	pagedID, err := s.storage.GetRelation(ctx, f, rel, limit)
	if err != nil {
		return nil, err
	}

	p := &def.PagedIDs{
		List:     pagedID.List,
		HasMore:  pagedID.HasMore,
		NextFrom: pagedID.NextFrom,
	}

	return p, nil
}

// GetByIds implements StorageServer method
func (s Wrapper) GetByIds(ctx context.Context, req *def.GetByIDsReq) (*def.EList, error) {
	es, err := s.storage.GetByIds(ctx, req.Type, req.GetIds())
	if err != nil {
		return nil, err
	}

	list := make([]*def.E, 0)
	for _, e := range es {
		list = append(list, toDef(e))
	}

	return &def.EList{
		List: list,
	}, nil
}

// GetByQuery implements StorageServer method
func (s Wrapper) GetByQuery(ctx context.Context, req *def.GetByQueryReq) (*def.Paged, error) {
	queries := make([]storage.Query, 0)
	for _, q := range req.Queries {
		_q := storage.Query{
			Field: q.Field,
			Op:    storage.Op(q.Op),
			Value: transValue(q.Value, q.ValueType),
		}

		queries = append(queries, _q)
	}

	sorts := make(map[string]storage.SortDir)
	for k, v := range req.Sorts {
		sorts[k] = storage.SortDir(v)
	}

	limit := toLimit(req.Limit)

	paged, err := s.storage.GetByQuery(ctx, req.Type, queries, sorts, limit)
	if err != nil {
		return nil, err
	}

	list := make([]*def.E, 0)
	for _, e := range paged.List {
		list = append(list, toDef(e))
	}

	return &def.Paged{
		List:    list,
		HasMore: paged.HasMore,
		//From: paged.N
	}, nil
}

func transValue(val string, t def.ValueType) interface{} {
	var v interface{}
	var err error
	switch t {
	case def.ValueType_String:
		return val
	case def.ValueType_Int:
		v, err = strconv.ParseInt(val, 10, 32)
	case def.ValueType_Int64:
		v, err = strconv.ParseInt(val, 10, 64)
	case def.ValueType_Bool:
		v, err = strconv.ParseBool(val)
	case def.ValueType_Double:
		v, err = strconv.ParseFloat(val, 64)
	case def.ValueType_Bytes:
		v = []byte(val)
	default:
		err = fmt.Errorf("bad value type: %v", t)
	}

	if err != nil {
		logging.WWarn("fail to convert value", "value", val, "type", t.String(), "err", err)
		v = val
	}

	return v
}

// UpdateContent implements StorageServer method
func (s Wrapper) UpdateContent(ctx context.Context, req *def.UpdateContentReq) (*def.Empty, error) {
	updates := make([]storage.Update, 0)
	for _, u := range req.Updates {
		updates = append(updates, storage.Update{
			Field:  u.Field,
			Action: storage.UpdateAction(u.Action),
			Value:  u.Value,
		})
	}
	err := s.storage.UpdateContent(ctx, req.Type, req.Id, updates)
	if err != nil {
		return nil, err
	}

	return &def.Empty{}, nil
}

// SetState implements StorageServer method
func (s Wrapper) SetState(ctx context.Context, req *def.SetStateReq) (*def.Empty, error) {
	err := s.storage.SetState(ctx, req.Type, req.Ids, storage.State(req.State))
	return &def.Empty{}, err
}

// Delete implements StorageServer method
func (s Wrapper) Delete(ctx context.Context, req *def.DeleteReq) (*def.Empty, error) {
	err := s.storage.Delete(ctx, req.Type, req.Ids)
	return &def.Empty{}, err
}

// GetCounts implements StorageServer method
func (s Wrapper) GetCounts(ctx context.Context, req *def.GetCountsReq) (*def.Counts, error) {
	cs, err := s.storage.GetCounts(ctx, req.Type, req.Id)
	if err != nil {
		return nil, err
	}

	_cs := &def.Counts{
		Counts: make(map[string]*def.CountByState),
	}

	for k, v := range cs {
		cbs := &def.CountByState{
			Counts: v,
		}

		_cs.Counts[k] = cbs
	}

	return _cs, nil
}

// HasRelations implements Storage method
func (s Wrapper) HasRelations(ctx context.Context, req *def.HasRelationsReq) (*def.HasRelations, error) {
	rels, err := s.storage.HasRelations(ctx, toModel(req.From), toModel(req.To), req.Relations)
	if err != nil {
		return nil, err
	}

	return &def.HasRelations{
		Relations: rels,
	}, nil
}

// SetMeta implements StorageServer method
func (s Wrapper) SetMeta(ctx context.Context, req *def.SetMetaReq) (*def.Empty, error) {
	err := s.storage.SetMeta(ctx, req.Type, req.Id, req.Meta)
	if err != nil {
		return nil, err
	}

	return &def.Empty{}, nil
}

// DeleteMeta implements StorageServer method
func (s Wrapper) DeleteMeta(ctx context.Context, req *def.DeleteMetaReq) (*def.Empty, error) {
	err := s.storage.DeleteMeta(ctx, req.Type, req.Id, req.Keys)
	if err != nil {
		return nil, err
	}

	return &def.Empty{}, nil
}

// AddTags implements StorageServer method
func (s Wrapper) AddTags(ctx context.Context, req *def.UpdateTagsReq) (*def.Empty, error) {
	err := s.storage.AddTags(ctx, req.Type, req.Id, req.Tags)
	if err != nil {
		return nil, err
	}

	return &def.Empty{}, nil
}

// DeleteTags implements StorageServer method
func (s Wrapper) DeleteTags(ctx context.Context, req *def.UpdateTagsReq) (*def.Empty, error) {
	err := s.storage.DeleteTags(ctx, req.Type, req.Id, req.Tags)
	if err != nil {
		return nil, err
	}

	return &def.Empty{}, nil
}

func toModel(e *def.E) *storage.E {
	return &storage.E{
		Type:      e.Type,
		ID:        e.ID,
		ID1:       e.ID1,
		ID2:       e.ID2,
		ID3:       e.ID3,
		CTime:     e.CTime,
		UTime:     e.UTime,
		State:     storage.State(e.State),
		Score:     e.Score,
		Score1:    e.Score1,
		Tags:      e.GetTags(),
		Meta:      e.GetMeta(),
		Content:   e.GetContent(),
		Resources: e.GetResources(),
		//ID4:     e.ID4,
	}
}

func toDef(e *storage.E) *def.E {
	if e == nil {
		return nil
	}
	return &def.E{
		Type:      e.Type,
		ID:        e.ID,
		ID1:       e.ID1,
		ID2:       e.ID2,
		ID3:       e.ID3,
		CTime:     e.CTime,
		UTime:     e.UTime,
		State:     def.State(e.State),
		Score:     e.Score,
		Score1:    e.Score1,
		Tags:      e.Tags,
		Meta:      e.Meta,
		Content:   e.Content,
		Resources: e.Resources,
		//ID4:     e.ID4,
	}
}

func toLimit(l *def.Limit) storage.Limit {
	if l == nil {
		l = &def.Limit{
			Limit: 20,
		}
	}
	limit := storage.Limit{
		Limit: int(l.Limit),
		From:  l.From,
	}

	return limit
}
