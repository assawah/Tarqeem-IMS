package utl

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	MaxNormalNameLength = 300
	MinNameLen          = 2
)

func EntString(n string) ent.Field { return field.String(n).MinLen(MinNameLen) }
func EntNumber(n string) ent.Field { return field.Int(n).Min(0) }
