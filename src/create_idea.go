package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	ideasPath = "./static/data/"
	noimgPath = "./static/img/logo.png"
)

func uploadLogo(dirPath string, r *http.Request) error {
	var fileBytes []byte
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("logoUpload")
	if err != nil || file == nil {
		fileBytes, err = ioutil.ReadFile(noimgPath)
		if err != nil {
			return err
		}
	} else {
		fileBytes, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}
	}
	//ext := "." + strings.Split(handler.Filename, ".")[1]
	//Create temp file and fill
	if err := ioutil.WriteFile(dirPath+"/logo", fileBytes, 0755); err != nil {
		return err
	}
	return nil
}

func saveInfo(dirPath string, r *http.Request) error {
	header := r.FormValue("headerInput")
	desc := r.FormValue("descriptionInput")
	author := r.FormValue("authorInput")
	tags := []string{"1", "2"}
	email := "test@mail.ru"
	newIdea := Idea{header, desc, author, email, tags}
	byteResult, err := json.Marshal(newIdea)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dirPath+"/data.json", byteResult, 0755); err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func addIdea(r *http.Request) error {

	files, err := ioutil.ReadDir(ideasPath)
	if err != nil {
		return err
	}
	folderName := ideasPath + strconv.Itoa(len(files))
	if err := os.Mkdir(folderName, 0755); err != nil {
		return err
	}
	if err := uploadLogo(folderName, r); err != nil {
		return err
	}
	if err := saveInfo(folderName, r); err != nil {
		return err
	}
	return nil
}
func idea_editor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := addIdea(r); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var tmpl = template.Must(template.ParseFiles(
		templatesPath+"create_idea.html",
		templatesPath+"create_idea_main.html",
		templatesPath+"create_idea_editor.html"))
	tmpl.ExecuteTemplate(w, "create_idea.html", nil)
}
