package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the Issue.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_path").Optional().Nillable(), // Store file path
		field.String("file_name").Optional().Nillable(),
		field.Int64("file_size").Optional().Default(0),
	}
}

// Edges of the Comment.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		// Edge to the Issue entity (one-to-many relationship)
		edge.From("issue", Issue.Type).
			Ref("files"),
	}
}
