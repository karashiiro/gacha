// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/karashiiro/gacha/ent/series"
)

// Series is the model entity for the Series schema.
type Series struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeriesQuery when eager-loading is set.
	Edges SeriesEdges `json:"edges"`
}

// SeriesEdges holds the relations/edges for other nodes in the graph.
type SeriesEdges struct {
	// Drops holds the value of the drops edge.
	Drops []*Drop `json:"drops,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DropsOrErr returns the Drops value or an error if the edge
// was not loaded in eager-loading.
func (e SeriesEdges) DropsOrErr() ([]*Drop, error) {
	if e.loadedTypes[0] {
		return e.Drops, nil
	}
	return nil, &NotLoadedError{edge: "drops"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Series) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case series.FieldID:
			values[i] = &sql.NullInt64{}
		case series.FieldName:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Series", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Series fields.
func (s *Series) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case series.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint32(value.Int64)
		case series.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		}
	}
	return nil
}

// QueryDrops queries the "drops" edge of the Series entity.
func (s *Series) QueryDrops() *DropQuery {
	return (&SeriesClient{config: s.config}).QueryDrops(s)
}

// Update returns a builder for updating this Series.
// Note that you need to call Series.Unwrap() before calling this method if this Series
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Series) Update() *SeriesUpdateOne {
	return (&SeriesClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Series entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Series) Unwrap() *Series {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Series is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Series) String() string {
	var builder strings.Builder
	builder.WriteString("Series(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// SeriesSlice is a parsable slice of Series.
type SeriesSlice []*Series

func (s SeriesSlice) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}