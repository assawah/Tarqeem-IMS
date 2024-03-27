package main

import (
	"log"
	"net/http"
)

func NewHandler(s string, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := executor.ExecuteTemplate(w, s, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error: "+err.Error(), 500)
		}

	}
}
