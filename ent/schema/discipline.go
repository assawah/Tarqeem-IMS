package schema

import (
	"entgo.io/ent"
	. "github.com/tarqeem/template/utl/ent"
)

// Discipline holds the schema definition for the Discipline entity.
type Discipline struct {
	ent.Schema
}

// Fields of the Discipline.
func (Discipline) Fields() []ent.Field {
	return []ent.Field{
		String("name"),
	}
}

// Edges of the Discipline.
func (Discipline) Edges() []ent.Edge {
	return nil
}
