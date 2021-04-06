package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Uint32("object_id"),
		field.Float32("rate"),
		field.Uint32("series_id"),
	}
}

// Edges of the Drop.
func (Drop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("series", Series.Type).
			Ref("drops").
			Field("series_id").
			Required().
			Unique(),
	}
}
