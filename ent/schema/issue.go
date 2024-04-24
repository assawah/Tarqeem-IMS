package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Issue holds the schema definition for the Issue entity.
type Issue struct {
	ent.Schema
}

// Fields of the Issue.
func (Issue) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description"),
		field.String("Creator").Default("guest"),
		field.Enum("status").Values("Pending", "Approved", "Declined").Default("Pending"),
		field.String("date").Default(time.Now().Format("02 Jan 2006")),
	}
}

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	return []ent.Edge{
		// Edge to the Project entity (one-to-many relationship)
		edge.From("project", Project.Type).
			Ref("issues").Unique(),
		// Edge to the Comment entity (one-to-many relationship)
		edge.To("comments", Comment.Type),
		edge.To("files", File.Type),
	}
}
