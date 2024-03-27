package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	. "github.com/tarqeem/template/utl/ent"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		String("name"),
		String("title"),
		String("owner"),
		field.Enum("type").Values(
			"Chemical manufacturing",
			"Stadium musuem",
			"Dam",
			"Metal refining/processing",
			"Oil exploration/production",
			"Oil refining",
			"Natural gas processing",
			"Highway",
			"Power generation",
			"Water/wastewater",
			"Consumer products manufacturing",
		),
		NonNegative("top_level_packages_number"),
		NonNegative("joint_venture_number"),
		NonNegative("dollar_value"),
		NonNegative("dollar_value"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
