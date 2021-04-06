package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Series holds the schema definition for the Series entity.
type Series struct {
	ent.Schema
}

// Fields of the Series.
func (Series) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.String("name"),
	}
}

// Edges of the Series.
func (Series) Edges() []ent.Edge {
	return nil
}
