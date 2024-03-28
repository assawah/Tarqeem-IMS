package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	. "github.com/tarqeem/template/utl/ent"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{String("name"), Passowrd(), Email(), Phone(),
		Created_at(), String("Organization"), String("Title"),
		field.Enum("type").Values("coordinator", "member"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	// return []ent.Edge{edge.To("coordinator", Coordinator.Type).Unique()}
	return nil
}
