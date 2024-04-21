// other
package utl

import (
	"errors"
	"unicode"
)

func ValidatePassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}

func ValidPassword(p string) (ok bool, err error) {

	sevenOrMore, number, upper, special := ValidatePassword(p)

	switch {
	case !sevenOrMore:
		errors.Join(err, ErrPassSevenOrMore)
	case !number:
		errors.Join(err, ErrNumber)
	case !upper:
		errors.Join(err, ErrUpper)
	case !special:
		errors.Join(err, ErrSpecial)
	default:
		ok = true
	}
	return
}

var (
	ErrPassSevenOrMore = errors.New("password must have seven or more letters")
	ErrNumber          = errors.New("password must have at least one number")
	ErrUpper           = errors.New("password must have at least one upercase letter")
	ErrSpecial         = errors.New("password must have at least one special character")
)
