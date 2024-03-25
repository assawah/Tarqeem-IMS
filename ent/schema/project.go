package schema

import (
	"entgo.io/ent"
	. "github.com/tarqeem/ims/utl"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		EntString("name"),
		EntString("title"),
		EntString("owner"),
		// EntNormalString("type"), TODO determine what's a type first.
		EntNumber("top_level_packages_number"),
		EntNumber("joint_venture_number"),
		EntNumber("dollar_value"),
		EntNumber("dollar_value"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
