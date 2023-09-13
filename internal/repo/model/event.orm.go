package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// EventsN is a Events object, all fields are nullable
type EventsN struct {
	original    *eventsOriginal
	eventsModel *EventsModel

	Id        null.Int    `json:"id"`
	EventType null.String `json:"event_type"`
	Payload   null.String `json:"payload"`
	Status    null.String `json:"status"`
	CreatedAt null.Time
	UpdatedAt null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *EventsN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for Events
func (inst *EventsN) SetModel(eventsModel *EventsModel) {
	inst.eventsModel = eventsModel
}

// eventsOriginal is an object which stores original Events from database
type eventsOriginal struct {
	Id        null.Int
	EventType null.String
	Payload   null.String
	Status    null.String
	CreatedAt null.Time
	UpdatedAt null.Time
}

// Staled identify whether the object has been modified
func (inst *EventsN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &eventsOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.EventType != inst.original.EventType {
			return true
		}
		if inst.Payload != inst.original.Payload {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "event_type":
				if inst.EventType != inst.original.EventType {
					return true
				}
			case "payload":
				if inst.Payload != inst.original.Payload {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *EventsN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &eventsOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.EventType != inst.original.EventType {
			kv["event_type"] = inst.EventType
		}
		if inst.Payload != inst.original.Payload {
			kv["payload"] = inst.Payload
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "event_type":
				if inst.EventType != inst.original.EventType {
					kv["event_type"] = inst.EventType
				}
			case "payload":
				if inst.Payload != inst.original.Payload {
					kv["payload"] = inst.Payload
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *EventsN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.eventsModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.eventsModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a events
func (inst *EventsN) Delete(ctx context.Context) error {
	if inst.eventsModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.eventsModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *EventsN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type eventsScope struct {
	name  string
	apply func(builder query.Condition)
}

var eventsGlobalScopes = make([]eventsScope, 0)
var eventsLocalScopes = make([]eventsScope, 0)

// AddGlobalScopeForEvents assign a global scope to a model
func AddGlobalScopeForEvents(name string, apply func(builder query.Condition)) {
	eventsGlobalScopes = append(eventsGlobalScopes, eventsScope{name: name, apply: apply})
}

// AddLocalScopeForEvents assign a local scope to a model
func AddLocalScopeForEvents(name string, apply func(builder query.Condition)) {
	eventsLocalScopes = append(eventsLocalScopes, eventsScope{name: name, apply: apply})
}

func (m *EventsModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range eventsGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range eventsLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *EventsModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *EventsModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type Events struct {
	Id        int64  `json:"id"`
	EventType string `json:"event_type"`
	Payload   string `json:"payload"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w Events) ToEventsN(allows ...string) EventsN {
	if len(allows) == 0 {
		return EventsN{

			Id:        null.IntFrom(int64(w.Id)),
			EventType: null.StringFrom(w.EventType),
			Payload:   null.StringFrom(w.Payload),
			Status:    null.StringFrom(w.Status),
			CreatedAt: null.TimeFrom(w.CreatedAt),
			UpdatedAt: null.TimeFrom(w.UpdatedAt),
		}
	}

	res := EventsN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "event_type":
			res.EventType = null.StringFrom(w.EventType)
		case "payload":
			res.Payload = null.StringFrom(w.Payload)
		case "status":
			res.Status = null.StringFrom(w.Status)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w Events) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *EventsN) ToEvents() Events {
	return Events{

		Id:        w.Id.Int64,
		EventType: w.EventType.String,
		Payload:   w.Payload.String,
		Status:    w.Status.String,
		CreatedAt: w.CreatedAt.Time,
		UpdatedAt: w.UpdatedAt.Time,
	}
}

// EventsModel is a model which encapsulates the operations of the object
type EventsModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var eventsTableName = "events"

// EventsTable return table name for Events
func EventsTable() string {
	return eventsTableName
}

const (
	FieldEventsId        = "id"
	FieldEventsEventType = "event_type"
	FieldEventsPayload   = "payload"
	FieldEventsStatus    = "status"
	FieldEventsCreatedAt = "created_at"
	FieldEventsUpdatedAt = "updated_at"
)

// EventsFields return all fields in Events model
func EventsFields() []string {
	return []string{
		"id",
		"event_type",
		"payload",
		"status",
		"created_at",
		"updated_at",
	}
}

func SetEventsTable(tableName string) {
	eventsTableName = tableName
}

// NewEventsModel create a EventsModel
func NewEventsModel(db query.Database) *EventsModel {
	return &EventsModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           eventsTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *EventsModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *EventsModel) clone() *EventsModel {
	return &EventsModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *EventsModel) WithoutGlobalScopes(names ...string) *EventsModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *EventsModel) WithLocalScopes(names ...string) *EventsModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *EventsModel) Condition(builder query.SQLBuilder) *EventsModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *EventsModel) Find(ctx context.Context, id int64) (*EventsN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *EventsModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *EventsModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *EventsModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]EventsN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *EventsModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]EventsN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"event_type",
			"payload",
			"status",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "event_type":
			selectFields = append(selectFields, f)
		case "payload":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*EventsN, []interface{}) {
		var eventsVar EventsN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &eventsVar.Id)
			case "event_type":
				scanFields = append(scanFields, &eventsVar.EventType)
			case "payload":
				scanFields = append(scanFields, &eventsVar.Payload)
			case "status":
				scanFields = append(scanFields, &eventsVar.Status)
			case "created_at":
				scanFields = append(scanFields, &eventsVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &eventsVar.UpdatedAt)
			}
		}

		return &eventsVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	eventss := make([]EventsN, 0)
	for rows.Next() {
		eventsReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		eventsReal.original = &eventsOriginal{}
		_ = query.Copy(eventsReal, eventsReal.original)

		eventsReal.SetModel(m)
		eventss = append(eventss, *eventsReal)
	}

	return eventss, nil
}

// First return first result for given query
func (m *EventsModel) First(ctx context.Context, builders ...query.SQLBuilder) (*EventsN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new events to database
func (m *EventsModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all eventss to database
func (m *EventsModel) SaveAll(ctx context.Context, eventss []EventsN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, events := range eventss {
		id, err := m.Save(ctx, events)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a events to database
func (m *EventsModel) Save(ctx context.Context, events EventsN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, events.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new events or update it when it has a id > 0
func (m *EventsModel) SaveOrUpdate(ctx context.Context, events EventsN, onlyFields ...string) (id int64, updated bool, err error) {
	if events.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, events.Id.Int64, events, onlyFields...)
		return events.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, events, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *EventsModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *EventsModel) Update(ctx context.Context, builder query.SQLBuilder, events EventsN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, events.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *EventsModel) UpdateById(ctx context.Context, id int64, events EventsN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, events.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *EventsModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *EventsModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}