package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Time("created_at").Default(time.Now),
		// Add other fields as necessary
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// Edge to the Issue entity (one-to-many relationship)
		edge.From("issue", Issue.Type).
			Ref("comments"),
	}
}
