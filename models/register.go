package models

type Register struct {
	Name         string `form:"name"`
	Email        string `form:"Email"`
	Title        string `form:"Title"`
	Phone        string `form:"Phone"`
	Password     string `form:"Password"`
	Organization string `form:"Organization"`
	Type         string `form:"Type"`
	PageTitle    string
	Err          error
}

// TODO Fix regex values in tmplt
