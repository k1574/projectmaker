package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type IdeaCards struct {
	UserName  string
	IdeaCards []IdeaCard
}
type IdeaCard struct {
	Id          string
	Header      string
	Description string
	Author      string
	Email       string
	Tags        []string
	Preview     string
}

/*
var tmpl = template.Must(template.ParseFiles(
	templatesPath + "index.html"))
*/
func readCard(filename string) (Idea, error) {
	filePath := ideasPath + filename + "/data.json"
	fileContent, err := os.Open(filePath)
	if err != nil {
		return Idea{"err", "err", "err", "err", []string{}}, err
	}
	byteResult, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return Idea{"err", "err", "err", "err", []string{}}, err
	}
	var data Idea
	defer fileContent.Close()
	if err := json.Unmarshal([]byte(byteResult), &data); err != nil {
		return Idea{"err", "err", "err", "err", []string{}}, err
	}
	return data, nil
}
func index(w http.ResponseWriter, r *http.Request) {
	ideaCards, err := readIdeasFromDB()
	if err != nil {
		fmt.Println(err)
	}
	user, err := checkSession(r)
	if err != nil {
		ideaCards.UserName = "Гость"
	} else {
		ideaCards.UserName = user.Name
	}
	var tmpl = template.Must(template.ParseFiles(
		templatesPath + "project-list.htm"))
	tmpl.ExecuteTemplate(w, "project-list.htm", ideaCards)
}
func readIdeasFromDB() (result IdeaCards, err error) {
	dirs, err := ioutil.ReadDir(ideasPath)
	if err != nil {
		return
	}
	for i, dir := range dirs {
		var idea Idea
		idea, err = readCard(dir.Name())
		if err != nil {
			return
		}
		var photoPath string
		if checkFileExists(ideasPath + dir.Name() + "/logo") {
			photoPath = ideasPath + dir.Name() + "/logo"
		} else {
			//photoPath = "./static/noimg.png"
			photoPath = ""
		}
		ideaCard := IdeaCard{strconv.Itoa(i), idea.Header, idea.Description, idea.Author, idea.Email, idea.Tags, photoPath}
		result.IdeaCards = append(result.IdeaCards, ideaCard)
	}
	return
}
func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}
