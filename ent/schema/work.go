package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Work holds the schema definition for the Work entity.
type Work struct {
	ent.Schema
}

// Fields of the Work.
func (Work) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255),
	}
}

// Edges of the Work.
func (Work) Edges() []ent.Edge {
	return nil
}
