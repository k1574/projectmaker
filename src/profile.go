package main

import (
	"fmt"
	"net/http"
	"errors"
)

func checkSession(r *http.Request) (*User, error) {
	session, err := cookie.Get(r, "user-session")
	var(
		email string
		password string
		user *User
		ok bool
	)
	email, ok = session.Values["email"].(string)
	if !ok {
		err = errors.New("Check session: Failed to read cookie")
		return user, err
	}
	password, ok = session.Values["password"].(string)
	if !ok {
		err = errors.New("Check session: Failed to read cookie")
		return user, err
	}
	user = GetUser(email)
	if user != nil && user.Password != password {
		user = nil
	}
	return user, err
}
func profile(w http.ResponseWriter, r *http.Request) {
	user, err := checkSession(r)

	if err != nil {
		http.Redirect(w, r, "login", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "Hello, "+user.Name+"!")
}
