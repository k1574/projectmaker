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

var (
	usersPath = "./users/"
	users     = []User{}
)

type User struct {
	Name     string
	Email    string
	Password string
}

func saveUser(dirPath string, r *http.Request) error {
	user := r.FormValue("user")
	email := r.FormValue("email")
	password := r.FormValue("password")

	newUser := User{user, email, password}
	users = append(users, newUser)
	byteResult, err := json.Marshal(newUser)
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

func addUser(r *http.Request) error {
	files, err := ioutil.ReadDir(usersPath)
	if err != nil {
		return err
	}
	email := r.FormValue("email")

	if GetUser(email) != nil {
		return errors.New("Add user: User Exists")
	}
	folderName := usersPath + strconv.Itoa(len(files))
	if err := os.Mkdir(folderName, 0755); err != nil {
		return err
	}
	if err := saveUser(folderName, r); err != nil {
		return err
	}
	return nil
}
func readUser(filename string) (User, error) {
	filePath := usersPath + filename + "/data.json"
	fileContent, err := os.Open(filePath)
	if err != nil {
		return User{"err", "err", "err"}, err
	}
	byteResult, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return User{"err", "err", "err"}, err
	}
	var data User
	defer fileContent.Close()
	if err := json.Unmarshal([]byte(byteResult), &data); err != nil {
		return User{"err", "err", "err"}, err
	}
	return data, nil
}
func readUsersFromDB() (result []User, err error) {
	dirs, err := ioutil.ReadDir(usersPath)
	if err != nil {
		return
	}
	for _, dir := range dirs {
		var user User
		user, err = readUser(dir.Name())
		if err != nil {
			return
		}
		result = append(result, user)
	}
	return
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if err := addUser(r); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		templatesPath + "sign-up.htm"))
	tmpl.ExecuteTemplate(w, "sign-up.htm", nil)
}
