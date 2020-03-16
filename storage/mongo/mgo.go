package mongo

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/butters-mars/tiki/logging"

	"github.com/butters-mars/taka/idgen"
	"github.com/butters-mars/taka/storage"
)

const (
	fieldCTime        = "Ct"
	fieldUTime        = "Ut"
	fieldID           = "Id"
	colCountsSuffix   = "counts"
	colRelationSuffix = "relation"
)

var (
	mut           = &sync.Mutex{}
	modifiedTypes = make(map[string]bool)
)

type storageImpl struct {
	session *mgo.Session
	idg     idgen.Service
}

// NewStorage -
func NewStorage(url string) (storage.Service, error) {
	s, err := mgo.Dial(url)
	if err != nil {
		logging.WError("fail to connect to mongo", "url", url, "err", err)
		return nil, err
	}

	_idgen, err := idgen.NewSnowFlakeIDGen(1)
	if err != nil {
		logging.WError("fail to create idgen", "err", err)
		return nil, err
	}

	return &storageImpl{
		session: s,
		idg:     _idgen,
	}, nil
}

func (st *storageImpl) Create(ctx context.Context, typ string, m storage.Model) error {
	ss := st.session.Copy()
	defer ss.Close()

	// t := reflect.TypeOf(m)
	// if t.Kind() == reflect.Ptr {
	// 	t = t.Elem()
	// }

	// typ := strings.ToLower(t.Name())
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// set ctime & id
	now := time.Now().UTC().Unix()
	ct := v.FieldByName(fieldCTime)
	if ct.IsValid() {
		ct.SetInt(now)
	}

	ut := v.FieldByName(fieldUTime)
	if ut.IsValid() {
		ut.SetInt(now)
	}

	id := v.FieldByName(fieldID)
	var _id int64
	if id.IsValid() {
		_id, _ = st.idg.GenID()
		logging.WDebug("set id", "id", _id)
		id.SetInt(_id)
	}

	col := ss.DB(typ).C(typ)

	logging.WDebug("do insert ...")
	err := col.Insert(m)
	if err != nil {
		logging.WError("fail to create entity in mongo", "type", typ, "err", err)
		return err
	}

	logging.WInfo("create entity in mongo OK", "type", typ, "id", _id)
	return err
}

func (st *storageImpl) Update(ctx context.Context, typ string, q storage.Query, updates []storage.Update) error {
	if q.Op != storage.Eq {
		return fmt.Errorf("only eq is supported when update")
	}

	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}

	// build update
	sel := bson.M{}
	sel[q.Field] = q.Value

	update := bson.M{}
	set := bson.M{
		"ut": time.Now().Unix(),
	}

	inc := bson.M{}
	for _, u := range updates {
		switch u.Action {
		case storage.Set:
			set[u.Field] = u.Value
		case storage.Incr:
			t := reflect.TypeOf(u.Value).Kind()
			if t != reflect.Int && t != reflect.Int64 {
				return fmt.Errorf("op inc with value not int/int64: %v", u.Value)
			}
			inc[u.Field] = u.Value
		case storage.Remove:
			return fmt.Errorf("op remove not supported yet when update")
		case storage.Add:
			return fmt.Errorf("op add not supported yet when update")
		}
	}

	if len(set) > 0 {
		update["$set"] = set
	}
	if len(inc) > 0 {
		update["$inc"] = set
	}

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(typ)
	err := col.Update(sel, update)
	if err != nil {
		logging.WError("fail to update obj", "type", typ, "query", q, "err", err)
		return err
	}

	logging.WDebug("update obj OK", "type", typ, "query", q)
	return nil
}

func (st *storageImpl) GetByID(ctx context.Context, typ string, id int64, out interface{}) error {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(typ)
	err := col.FindId(id).One(out)
	if err != nil {
		logging.WError("getByID find obj", "type", typ, "id", id, "err", err)
		return err
	}

	logging.WDebug("getByID OK", "type", typ, "id", id)
	return nil
}

func (st *storageImpl) GetByIDs(ctx context.Context, typ string, ids []int64, out interface{}) (err error) {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}

	defer func() {
		if r := recover(); r != nil {
			logging.WError("panic when getByIDs", "err", r)
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(typ)
	err = col.Find(bson.M{"_id": bson.M{"$in": ids}}).All(out)
	if err != nil {
		logging.WError("fail to getByIDs", "type", typ, "ids", ids, "err", err)
		return err
	}

	logging.WDebug("getByIDs OK", "type", typ, "ids", ids)
	return nil
}

func (st *storageImpl) GetOneByQuery(ctx context.Context, typ string, qs []storage.Query, obj interface{}) error {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}

	qry := buildQuery(qs)
	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(typ)
	err := col.Find(qry).One(obj)
	if err != nil {
		logging.WError("fail to getOneByQuery", "type", typ, "query", qry, "err", err)
		return err
	}

	logging.WDebug("getOneByQuery OK", "type", typ, "query", qry)
	return nil
}

func (st *storageImpl) GetByQuery(ctx context.Context, typ string, qs []storage.Query, sorts []storage.Sort, limit storage.Limit, slice interface{}) (bool, error) {
	typ = strings.ToLower(typ)
	if typ == "" {
		return false, fmt.Errorf("type not specified")
	}

	qry := buildQuery(qs)
	srt := buildSort(sorts)

	ss := st.session.Copy()
	defer ss.Close()

	n := limit.Limit
	if n == 0 {
		n = 20
		limit.Limit = n
	}

	// TODO do not use skip
	col := ss.DB(typ).C(typ)
	q := col.Find(qry).Sort(srt...).Limit(n)
	if limit.Page > 0 {
		logging.WWarn("query using SKIP", "type", typ, "query", qry)
		q = q.Skip(n * limit.Page)
	}
	err := q.All(slice)
	if err != nil {
		logging.WError("fail to getByQuery", "type", typ, "query", qry,
			"sort", srt, "limit", limit, "err", err)
		return false, err
	}

	len := reflect.ValueOf(slice).Elem().Len()
	more := len == n

	logging.WDebug("getByQuery OK", "type", typ, "query", qry,
		"sort", srt, "limit", limit, "more", more)
	return more, nil
}

func (st *storageImpl) SetState(ctx context.Context, typ string, id int64, state int32) error {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(typ)
	err := col.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"state": state}})
	if err != nil {
		logging.WError("fail to setState", "type", typ, "id", id, "state", state, "err", err)
		return err
	}

	logging.WDebug("setState OK", "type", typ, "id", id, "state", state)
	return nil
}

func (st *storageImpl) GetCount(ctx context.Context, typ string, id int64) (map[string]int64, error) {
	typ = strings.ToLower(typ)
	if typ == "" {
		return nil, fmt.Errorf("type not specified")
	}

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(fmt.Sprintf("%s_%s", typ, colCountsSuffix))
	cs := make(map[string]int64)
	err := col.Find(bson.M{"id": id}).One(&cs)
	if err != nil {
		if err == mgo.ErrNotFound {
			return cs, nil
		}

		logging.WError("fail to get count", "type", typ, "id", id, "err", err)
		return nil, err
	}

	logging.WDebug("get count OK", "type", typ, "id", id)
	return cs, nil
}

func (st *storageImpl) IncrCount(ctx context.Context, typ string, id int64, delta map[string]int64) error {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}

	ss := st.session.Copy()
	defer ss.Close()

	return incrCount(ss, typ, id, delta)
}

func incrCount(ss *mgo.Session, typ string, id int64, delta map[string]int64) error {
	col := ss.DB(typ).C(fmt.Sprintf("%s_%s", typ, colCountsSuffix))
	up := bson.M{}
	for k, v := range delta {
		up[k] = v
	}

	err := col.Update(bson.M{"id": id}, bson.M{"$inc": up})
	if err != nil {
		if err != mgo.ErrNotFound {
			logging.WError("fail to incr count", "type", typ, "id", id, "delta", delta, "err", err)
			return err
		}

		delta["id"] = id
		err = col.Insert(delta)
		if err != nil {
			logging.WError("fail to incr count (insert new)", "type", typ, "id", id, "delta", delta, "err", err)
			return err
		}
	}

	logging.WDebug("incr count OK", "type", typ, "id", id, "delta", delta)
	return nil
}

type relation struct {
	ID int64 `bson:"id"`
	To int64 `bson:"to"`
	CT int64 `bson:"ct"`
}

func (st *storageImpl) AddRelation(ctx context.Context, typ string, id, to int64, rel string) error {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}
	if rel == "" {
		return fmt.Errorf("relation not specified")
	}

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(fmt.Sprintf("%s_%s", typ, rel))
	err := col.Find(bson.M{"id": id, "to": to}).One(&relation{})
	if err == nil {
		return fmt.Errorf("relation %s already exists", rel)
	}
	if err != nil && err != mgo.ErrNotFound {
		logging.WError("fail to check relation existence when adding", "type", typ, "id", id, "to", to, "rel", rel, "err", err)
		return err
	}

	err = col.Insert(&relation{
		ID: id,
		To: to,
		CT: time.Now().Unix(),
	})
	if err != nil {
		logging.WError("fail to add relation", "type", typ, "id", id, "to", to, "rel", rel, "err", err)
		return err
	}

	delta := make(map[string]int64)
	delta[rel] = 1
	err = incrCount(ss, typ, id, delta)
	if err != nil {
		logging.WError("fail to incr count when add relation", "type", typ, "id", id, "to", to, "rel", rel, "err", err)
	}

	logging.WDebug("add relation OK", "type", typ, "id", id, "to", to, "rel", rel)
	return nil
}

func (st *storageImpl) RemoveRelation(ctx context.Context, typ string, id, to int64, rel string) error {
	typ = strings.ToLower(typ)
	if typ == "" {
		return fmt.Errorf("type not specified")
	}
	if rel == "" {
		return fmt.Errorf("relation not specified")
	}

	ss := st.session.Copy()
	defer ss.Close()

	col := ss.DB(typ).C(fmt.Sprintf("%s_%s", typ, rel))
	err := col.Find(bson.M{"id": id, "to": to}).One(&relation{})
	if err != nil {
		if err == mgo.ErrNotFound {
			return err
		}

		logging.WError("fail to check existence when removing", "typ", typ, "id", id, "to", to, "rel", rel, "err", err)
		return err
	}

	err = col.Remove(bson.M{
		"id": id,
		"to": to,
	})
	if err != nil {
		logging.WError("fail to remove relation", "type", typ, "id", id, "to", to, "rel", rel, "err", err)
		return err
	}

	delta := make(map[string]int64)
	delta[rel] = -1
	err = incrCount(ss, typ, id, delta)
	if err != nil {
		logging.WError("fail to desc count when remove relation", "type", typ, "id", id, "to", to, "rel", rel, "err", err)
	}

	logging.WDebug("remove relation OK", "type", typ, "id", id, "to", to, "rel", rel)
	return nil
}

func (st *storageImpl) GetRelated(ctx context.Context, typ string, id int64, rel string, limit storage.Limit) ([]int64, error) {
	return nil, nil
}

func buildSort(ss []storage.Sort) []string {
	var _ss []string
	for _, s := range ss {
		_s := s.Field
		if s.Dir == storage.Desc {
			_s = "-" + _s
		}
		_ss = append(_ss, _s)
	}

	return _ss
}

func buildQuery(qs []storage.Query) bson.M {
	qry := bson.M{}
	for _, q := range qs {
		val := q.Value
		switch q.Op {
		case storage.Eq:
			break
		case storage.Ne:
			val = bson.M{"$ne": val}
		case storage.Lt:
			val = bson.M{"$lt": val}
		case storage.Le:
			val = bson.M{"$lte": val}
		case storage.Gt:
			val = bson.M{"$gt": val}
		case storage.Ge:
			val = bson.M{"$gte": val}
		case storage.In:
			val = bson.M{"$in": val}
		}

		qry[q.Field] = val
	}

	return qry
}
