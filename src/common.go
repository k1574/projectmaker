package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	templatesPath = "static/html/"
)

type Idea struct {
	Header      string
	Description string
	Author      string
	Email       string
	Tags        []string
}

func getCards() IdeaCards {
	ideaCards := IdeaCards{}

	dirs, err := ioutil.ReadDir(ideasPath)
	if err != nil {
		fmt.Println(err)
	}
	for i, dir := range dirs {
		idea, err := readCard(dir.Name())
		if err != nil {
			fmt.Println(err)
		}
		var photoPath string
		if checkFileExists(ideasPath + dir.Name() + "/logo") {
			photoPath = ideasPath + dir.Name() + "/logo"
		} else {
			photoPath = "/static/noimg.png"
		}
		ideaCard := IdeaCard{strconv.Itoa(i), idea.Header, idea.Description, idea.Author, idea.Email, idea.Tags, photoPath}
		ideaCards.IdeaCards = append(ideaCards.IdeaCards, ideaCard)
	}
	return ideaCards
}
func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" || strings.HasSuffix(r.URL.Path, ".json") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
