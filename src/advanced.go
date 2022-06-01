package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ()

func advanced(w http.ResponseWriter, r *http.Request) {
	cards := getCards()
	vars := mux.Vars(r)["id"]
	var tmpl = template.Must(template.ParseFiles(
		templatesPath + "advanced.html"))
	id, err := strconv.Atoi(vars)
	if err != nil || id > len(cards.IdeaCards) {
		id = 0
	}
	tmpl.ExecuteTemplate(w, "advanced.html", nil) //cards.IdeaCards[id].Markdown)
}
