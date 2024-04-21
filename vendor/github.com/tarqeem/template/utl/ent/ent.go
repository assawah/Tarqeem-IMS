package ent

import (
	"fmt"
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/tarqeem/template/utl"
)

var (
	MaxNormalNameLength = 300
	MinNameLen          = 2
	MinPhoneLen         = 9
	EmailRegex          = regexp.MustCompile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
	PhoneRegex          = regexp.MustCompile(`^\+?\d+$`)
	CreatedTableName    = "created_at"
	PasswordTableName   = "passowrd"
	PhoneTableName      = "phone"
)

// return a field `created_at`. Configure name with CreatedTableName
func Created_at() ent.Field { return field.Time(CreatedTableName).Default(time.Now) }

// return a string field with `n` name and `MinNameLen` minimum length
func String(n string) ent.Field { return field.String(n).MinLen(MinNameLen) }

// return a string field with `n` name and `MinNameLen` minimum length, must
// be one of the ones in `l`
func StringOneOf(n string, l []string) ent.Field {
	return field.String(n).Validate(func(s string) error {
		for _, v := range l {
			if s == v {
				return nil
			}
		}
		return fmt.Errorf("No such a type: %s", s)
	})
}

// return a field `password`. Configure name with `CreatedTableName` and Regex
// with PassRegex
func Passowrd() ent.Field {
	return field.String(PasswordTableName).Sensitive().NotEmpty().
		Validate(func(s string) error {
			_, err := utl.ValidPassword(s)
			return err
		})
}

// return a non negative integer field
func NonNegative(n string) ent.Field { return field.Int(n).Min(0) }

// return a field named `phone`. Configure name with `PhoneTableName` and Regex
// with `PhoneRegex`
func Phone() ent.Field {
	return field.String(PhoneTableName).MinLen(MinPhoneLen).NotEmpty().Match(PhoneRegex).Unique()
}

// return a field named `email`. Configure name with `EmailTableName` and Regex
// with `EmailRegex`
func Email() ent.Field {
	return field.String("email").MaxLen(MaxNormalNameLength).NotEmpty().Match(EmailRegex)
}
