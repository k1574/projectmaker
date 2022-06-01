package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	cookie *sessions.CookieStore
)

func setupRouters(rtr *mux.Router) {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", neuter(fs)))

	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/login", login).Methods("GET")
	rtr.HandleFunc("/sign-up", signUp).Methods("GET")
	rtr.HandleFunc("/autorize", Authorization).Methods("GET")
	rtr.HandleFunc("/profile", profile).Methods("GET")
	rtr.HandleFunc("/createUser", createUser).Methods("GET")
	rtr.HandleFunc("/create_idea", idea_editor).Methods("POST", "GET")
	rtr.HandleFunc("/advanced{id:[0-9]+}", advanced).Methods("GET")
	http.Handle("/", rtr)
}
func initCookie() {
	cookie = sessions.NewCookieStore([]byte("Users"))
}
func loadUsersData() (err error) {
	users, err = readUsersFromDB()
	return err
}

func main() {
	rtr := mux.NewRouter()
	setupRouters(rtr)
	if err := loadUsersData(); err != nil {
		log.Fatal(err)
	}
	initCookie()
	fmt.Println("Starting")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
