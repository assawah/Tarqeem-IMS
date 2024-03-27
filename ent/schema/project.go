package schema

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	. "github.com/tarqeem/template/utl/ent"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

var ValidProjectDeliverayStratigies []string
var ValidProjectContractingStratigies []string
var ValidProjectTypes []string = []string{
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
}

var ValidProjectStages []string = []string{
	"Front-end planning",
	"Design",
	"Procurement",
	"Construction",
	"Commissions",
	"Start-up",
	"Completed",
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		String("name"),
		String("owner"),
		String("location"),
		field.String("type").Validate(func(s string) error {
			for _, v := range ValidProjectTypes {
				if s == v {
					return nil
				}
			}
			return fmt.Errorf("No such a type: %s", s)
		}),
		field.Enum("Project_nature"),
		NonNegative("top_level_packages_number"),
		NonNegative("joint_venture_number"),
		String("execution_location"),
		NonNegative("involved_stockholders"),
		NonNegative("dollar_value"),

		field.String("stage").Validate(func(s string) error {
			for _, v := range ValidProjectStages {
				if s == v {
					return nil
				}
			}
			return fmt.Errorf("No such a type: %s", s)
		}),

		field.String("delivery_stratigies").Validate(func(s string) error {
			for _, v := range ValidProjectDeliverayStratigies {
				if s == v {
					return nil
				}
			}
			return fmt.Errorf("No such a type: %s", s)
		}),

		field.String("contracting_stratigies").Validate(func(s string) error {
			for _, v := range ValidProjectContractingStratigies {
				if s == v {
					return nil
				}
			}
			return fmt.Errorf("No such a type: %s", s)
		}),
	}
}

// Edges of the Project.
// TODO
func (Project) Edges() []ent.Edge {
	return nil
}
