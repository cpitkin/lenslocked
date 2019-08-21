package controllers

import (
	"net/http"

	schema "github.com/gorilla/Schema"
)

// ParseForm will parse form data and return the result to the specified struct
func ParseForm(r *http.Request, dst interface{}) error {
	var err error

	err = r.ParseForm()

	decoder := schema.NewDecoder()
	err = decoder.Decode(dst, r.PostForm)

	if err != nil {
		return err
	}

	return nil
}
