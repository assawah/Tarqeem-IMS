package schema

import (
	. "github.com/tarqeem/ims/translate"

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
		String("owner"),
		String("location"),
		StringOneOf("type", ValidProjectTypes),
		field.Enum("Project_nature"),
		NonNegative("top_level_packages_number"),
		NonNegative("joint_venture_number"),
		String("execution_location"),
		NonNegative("involved_stockholders"),
		NonNegative("dollar_value"),
		StringOneOf("stage", ValidProjectStages),
		StringOneOf("delivery_stratigies", ValidProjectDeliverayStratigies),
		StringOneOf("contracting_stratigies", ValidProjectContractingStratigies),
	}
}

// Edges of the Project.
// TODO
func (Project) Edges() []ent.Edge {
	return nil
}
