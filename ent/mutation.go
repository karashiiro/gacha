// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/karashiiro/gacha/ent/drop"
	"github.com/karashiiro/gacha/ent/predicate"
	"github.com/karashiiro/gacha/ent/series"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDrop   = "Drop"
	TypeSeries = "Series"
)

// DropMutation represents an operation that mutates the Drop nodes in the graph.
type DropMutation struct {
	config
	op            Op
	typ           string
	id            *uint32
	object_id     *uint32
	addobject_id  *uint32
	rate          *float32
	addrate       *float32
	clearedFields map[string]struct{}
	series        *uint32
	clearedseries bool
	done          bool
	oldValue      func(context.Context) (*Drop, error)
	predicates    []predicate.Drop
}

var _ ent.Mutation = (*DropMutation)(nil)

// dropOption allows management of the mutation configuration using functional options.
type dropOption func(*DropMutation)

// newDropMutation creates new mutation for the Drop entity.
func newDropMutation(c config, op Op, opts ...dropOption) *DropMutation {
	m := &DropMutation{
		config:        c,
		op:            op,
		typ:           TypeDrop,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDropID sets the ID field of the mutation.
func withDropID(id uint32) dropOption {
	return func(m *DropMutation) {
		var (
			err   error
			once  sync.Once
			value *Drop
		)
		m.oldValue = func(ctx context.Context) (*Drop, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Drop.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDrop sets the old Drop of the mutation.
func withDrop(node *Drop) dropOption {
	return func(m *DropMutation) {
		m.oldValue = func(context.Context) (*Drop, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DropMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DropMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Drop entities.
func (m *DropMutation) SetID(id uint32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *DropMutation) ID() (id uint32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetObjectID sets the "object_id" field.
func (m *DropMutation) SetObjectID(u uint32) {
	m.object_id = &u
	m.addobject_id = nil
}

// ObjectID returns the value of the "object_id" field in the mutation.
func (m *DropMutation) ObjectID() (r uint32, exists bool) {
	v := m.object_id
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectID returns the old "object_id" field's value of the Drop entity.
// If the Drop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DropMutation) OldObjectID(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldObjectID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldObjectID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectID: %w", err)
	}
	return oldValue.ObjectID, nil
}

// AddObjectID adds u to the "object_id" field.
func (m *DropMutation) AddObjectID(u uint32) {
	if m.addobject_id != nil {
		*m.addobject_id += u
	} else {
		m.addobject_id = &u
	}
}

// AddedObjectID returns the value that was added to the "object_id" field in this mutation.
func (m *DropMutation) AddedObjectID() (r uint32, exists bool) {
	v := m.addobject_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetObjectID resets all changes to the "object_id" field.
func (m *DropMutation) ResetObjectID() {
	m.object_id = nil
	m.addobject_id = nil
}

// SetRate sets the "rate" field.
func (m *DropMutation) SetRate(f float32) {
	m.rate = &f
	m.addrate = nil
}

// Rate returns the value of the "rate" field in the mutation.
func (m *DropMutation) Rate() (r float32, exists bool) {
	v := m.rate
	if v == nil {
		return
	}
	return *v, true
}

// OldRate returns the old "rate" field's value of the Drop entity.
// If the Drop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DropMutation) OldRate(ctx context.Context) (v float32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRate: %w", err)
	}
	return oldValue.Rate, nil
}

// AddRate adds f to the "rate" field.
func (m *DropMutation) AddRate(f float32) {
	if m.addrate != nil {
		*m.addrate += f
	} else {
		m.addrate = &f
	}
}

// AddedRate returns the value that was added to the "rate" field in this mutation.
func (m *DropMutation) AddedRate() (r float32, exists bool) {
	v := m.addrate
	if v == nil {
		return
	}
	return *v, true
}

// ResetRate resets all changes to the "rate" field.
func (m *DropMutation) ResetRate() {
	m.rate = nil
	m.addrate = nil
}

// SetSeriesID sets the "series_id" field.
func (m *DropMutation) SetSeriesID(u uint32) {
	m.series = &u
}

// SeriesID returns the value of the "series_id" field in the mutation.
func (m *DropMutation) SeriesID() (r uint32, exists bool) {
	v := m.series
	if v == nil {
		return
	}
	return *v, true
}

// OldSeriesID returns the old "series_id" field's value of the Drop entity.
// If the Drop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DropMutation) OldSeriesID(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSeriesID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSeriesID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSeriesID: %w", err)
	}
	return oldValue.SeriesID, nil
}

// ResetSeriesID resets all changes to the "series_id" field.
func (m *DropMutation) ResetSeriesID() {
	m.series = nil
}

// ClearSeries clears the "series" edge to the Series entity.
func (m *DropMutation) ClearSeries() {
	m.clearedseries = true
}

// SeriesCleared returns if the "series" edge to the Series entity was cleared.
func (m *DropMutation) SeriesCleared() bool {
	return m.clearedseries
}

// SeriesIDs returns the "series" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// SeriesID instead. It exists only for internal usage by the builders.
func (m *DropMutation) SeriesIDs() (ids []uint32) {
	if id := m.series; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetSeries resets all changes to the "series" edge.
func (m *DropMutation) ResetSeries() {
	m.series = nil
	m.clearedseries = false
}

// Op returns the operation name.
func (m *DropMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Drop).
func (m *DropMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DropMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.object_id != nil {
		fields = append(fields, drop.FieldObjectID)
	}
	if m.rate != nil {
		fields = append(fields, drop.FieldRate)
	}
	if m.series != nil {
		fields = append(fields, drop.FieldSeriesID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DropMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case drop.FieldObjectID:
		return m.ObjectID()
	case drop.FieldRate:
		return m.Rate()
	case drop.FieldSeriesID:
		return m.SeriesID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DropMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case drop.FieldObjectID:
		return m.OldObjectID(ctx)
	case drop.FieldRate:
		return m.OldRate(ctx)
	case drop.FieldSeriesID:
		return m.OldSeriesID(ctx)
	}
	return nil, fmt.Errorf("unknown Drop field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DropMutation) SetField(name string, value ent.Value) error {
	switch name {
	case drop.FieldObjectID:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectID(v)
		return nil
	case drop.FieldRate:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRate(v)
		return nil
	case drop.FieldSeriesID:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSeriesID(v)
		return nil
	}
	return fmt.Errorf("unknown Drop field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DropMutation) AddedFields() []string {
	var fields []string
	if m.addobject_id != nil {
		fields = append(fields, drop.FieldObjectID)
	}
	if m.addrate != nil {
		fields = append(fields, drop.FieldRate)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DropMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case drop.FieldObjectID:
		return m.AddedObjectID()
	case drop.FieldRate:
		return m.AddedRate()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DropMutation) AddField(name string, value ent.Value) error {
	switch name {
	case drop.FieldObjectID:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddObjectID(v)
		return nil
	case drop.FieldRate:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRate(v)
		return nil
	}
	return fmt.Errorf("unknown Drop numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DropMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DropMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DropMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Drop nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DropMutation) ResetField(name string) error {
	switch name {
	case drop.FieldObjectID:
		m.ResetObjectID()
		return nil
	case drop.FieldRate:
		m.ResetRate()
		return nil
	case drop.FieldSeriesID:
		m.ResetSeriesID()
		return nil
	}
	return fmt.Errorf("unknown Drop field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DropMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.series != nil {
		edges = append(edges, drop.EdgeSeries)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DropMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case drop.EdgeSeries:
		if id := m.series; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DropMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DropMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DropMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedseries {
		edges = append(edges, drop.EdgeSeries)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DropMutation) EdgeCleared(name string) bool {
	switch name {
	case drop.EdgeSeries:
		return m.clearedseries
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DropMutation) ClearEdge(name string) error {
	switch name {
	case drop.EdgeSeries:
		m.ClearSeries()
		return nil
	}
	return fmt.Errorf("unknown Drop unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DropMutation) ResetEdge(name string) error {
	switch name {
	case drop.EdgeSeries:
		m.ResetSeries()
		return nil
	}
	return fmt.Errorf("unknown Drop edge %s", name)
}

// SeriesMutation represents an operation that mutates the Series nodes in the graph.
type SeriesMutation struct {
	config
	op            Op
	typ           string
	id            *uint32
	name          *string
	clearedFields map[string]struct{}
	drops         map[uint32]struct{}
	removeddrops  map[uint32]struct{}
	cleareddrops  bool
	done          bool
	oldValue      func(context.Context) (*Series, error)
	predicates    []predicate.Series
}

var _ ent.Mutation = (*SeriesMutation)(nil)

// seriesOption allows management of the mutation configuration using functional options.
type seriesOption func(*SeriesMutation)

// newSeriesMutation creates new mutation for the Series entity.
func newSeriesMutation(c config, op Op, opts ...seriesOption) *SeriesMutation {
	m := &SeriesMutation{
		config:        c,
		op:            op,
		typ:           TypeSeries,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSeriesID sets the ID field of the mutation.
func withSeriesID(id uint32) seriesOption {
	return func(m *SeriesMutation) {
		var (
			err   error
			once  sync.Once
			value *Series
		)
		m.oldValue = func(ctx context.Context) (*Series, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Series.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSeries sets the old Series of the mutation.
func withSeries(node *Series) seriesOption {
	return func(m *SeriesMutation) {
		m.oldValue = func(context.Context) (*Series, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SeriesMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SeriesMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Series entities.
func (m *SeriesMutation) SetID(id uint32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *SeriesMutation) ID() (id uint32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *SeriesMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *SeriesMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Series entity.
// If the Series object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SeriesMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *SeriesMutation) ResetName() {
	m.name = nil
}

// AddDropIDs adds the "drops" edge to the Drop entity by ids.
func (m *SeriesMutation) AddDropIDs(ids ...uint32) {
	if m.drops == nil {
		m.drops = make(map[uint32]struct{})
	}
	for i := range ids {
		m.drops[ids[i]] = struct{}{}
	}
}

// ClearDrops clears the "drops" edge to the Drop entity.
func (m *SeriesMutation) ClearDrops() {
	m.cleareddrops = true
}

// DropsCleared returns if the "drops" edge to the Drop entity was cleared.
func (m *SeriesMutation) DropsCleared() bool {
	return m.cleareddrops
}

// RemoveDropIDs removes the "drops" edge to the Drop entity by IDs.
func (m *SeriesMutation) RemoveDropIDs(ids ...uint32) {
	if m.removeddrops == nil {
		m.removeddrops = make(map[uint32]struct{})
	}
	for i := range ids {
		m.removeddrops[ids[i]] = struct{}{}
	}
}

// RemovedDrops returns the removed IDs of the "drops" edge to the Drop entity.
func (m *SeriesMutation) RemovedDropsIDs() (ids []uint32) {
	for id := range m.removeddrops {
		ids = append(ids, id)
	}
	return
}

// DropsIDs returns the "drops" edge IDs in the mutation.
func (m *SeriesMutation) DropsIDs() (ids []uint32) {
	for id := range m.drops {
		ids = append(ids, id)
	}
	return
}

// ResetDrops resets all changes to the "drops" edge.
func (m *SeriesMutation) ResetDrops() {
	m.drops = nil
	m.cleareddrops = false
	m.removeddrops = nil
}

// Op returns the operation name.
func (m *SeriesMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Series).
func (m *SeriesMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SeriesMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, series.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SeriesMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case series.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SeriesMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case series.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Series field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SeriesMutation) SetField(name string, value ent.Value) error {
	switch name {
	case series.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Series field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SeriesMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SeriesMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SeriesMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Series numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SeriesMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SeriesMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SeriesMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Series nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SeriesMutation) ResetField(name string) error {
	switch name {
	case series.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Series field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SeriesMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.drops != nil {
		edges = append(edges, series.EdgeDrops)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SeriesMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case series.EdgeDrops:
		ids := make([]ent.Value, 0, len(m.drops))
		for id := range m.drops {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SeriesMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeddrops != nil {
		edges = append(edges, series.EdgeDrops)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SeriesMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case series.EdgeDrops:
		ids := make([]ent.Value, 0, len(m.removeddrops))
		for id := range m.removeddrops {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SeriesMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareddrops {
		edges = append(edges, series.EdgeDrops)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SeriesMutation) EdgeCleared(name string) bool {
	switch name {
	case series.EdgeDrops:
		return m.cleareddrops
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SeriesMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Series unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SeriesMutation) ResetEdge(name string) error {
	switch name {
	case series.EdgeDrops:
		m.ResetDrops()
		return nil
	}
	return fmt.Errorf("unknown Series edge %s", name)
}
