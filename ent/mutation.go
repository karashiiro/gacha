// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/karashiiro/gacha/ent/drop"
	"github.com/karashiiro/gacha/ent/predicate"

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
	TypeDrop = "Drop"
)

// DropMutation represents an operation that mutates the Drop nodes in the graph.
type DropMutation struct {
	config
	op            Op
	typ           string
	id            *uint32
	rate          *float32
	addrate       *float32
	series        *string
	clearedFields map[string]struct{}
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

// SetSeries sets the "series" field.
func (m *DropMutation) SetSeries(s string) {
	m.series = &s
}

// Series returns the value of the "series" field in the mutation.
func (m *DropMutation) Series() (r string, exists bool) {
	v := m.series
	if v == nil {
		return
	}
	return *v, true
}

// OldSeries returns the old "series" field's value of the Drop entity.
// If the Drop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DropMutation) OldSeries(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSeries is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSeries requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSeries: %w", err)
	}
	return oldValue.Series, nil
}

// ResetSeries resets all changes to the "series" field.
func (m *DropMutation) ResetSeries() {
	m.series = nil
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
	fields := make([]string, 0, 2)
	if m.rate != nil {
		fields = append(fields, drop.FieldRate)
	}
	if m.series != nil {
		fields = append(fields, drop.FieldSeries)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DropMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case drop.FieldRate:
		return m.Rate()
	case drop.FieldSeries:
		return m.Series()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DropMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case drop.FieldRate:
		return m.OldRate(ctx)
	case drop.FieldSeries:
		return m.OldSeries(ctx)
	}
	return nil, fmt.Errorf("unknown Drop field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DropMutation) SetField(name string, value ent.Value) error {
	switch name {
	case drop.FieldRate:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRate(v)
		return nil
	case drop.FieldSeries:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSeries(v)
		return nil
	}
	return fmt.Errorf("unknown Drop field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DropMutation) AddedFields() []string {
	var fields []string
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
	case drop.FieldRate:
		m.ResetRate()
		return nil
	case drop.FieldSeries:
		m.ResetSeries()
		return nil
	}
	return fmt.Errorf("unknown Drop field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DropMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DropMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DropMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DropMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DropMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DropMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DropMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Drop unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DropMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Drop edge %s", name)
}
