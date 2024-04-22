package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	. "github.com/tarqeem/template/utl/ent"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("password").Optional(),
		field.String("email").Unique(),
		field.String("phone").Optional(),
		Created_at(),
		field.String("organization").Optional(),
		field.String("title").Optional(),
		field.Bool("is_active").Default(true),
		field.Enum("type").Values("coordinator", "member"),
	}
}

// Edges of the Member.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("projects", Project.Type).
			Ref("members"),
		edge.To("leader_of_project", Project.Type),
		edge.To("coordinator_of_project", Project.Type),
	}
}
