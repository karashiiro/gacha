package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Drop holds the schema definition for the Drop entity.
type Drop struct {
	ent.Schema
}

// Fields of the Drop.
func (Drop) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Float32("rate"),
		field.Uint32("series"),
	}
}

// Edges of the Drop.
func (Drop) Edges() []ent.Edge {
	return nil
}
