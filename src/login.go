package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func IsValidUser(email string, password string) bool {
	user := GetUser(email)
	return (user != nil && user.Password == password)
}
func GetUser(email string) *User {
	for _, user := range users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
func Authorization(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if IsValidUser(email, password) {
		session, _ := cookie.Get(r, "user-session")
		session.Values["email"] = email
		session.Values["password"] = password
		session.Save(r, w)
		//fmt.Fprintf(w, "Successfully Logged In")
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		fmt.Fprintln(w, "Invalid login or password")
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	_, err := checkSession(r)
	if err == nil {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	var tmpl = template.Must(template.ParseFiles(
		templatesPath + "login.htm"))
	tmpl.ExecuteTemplate(w, "login.htm", nil)
}
