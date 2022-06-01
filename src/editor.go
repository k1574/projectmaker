package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var header, _ = ioutil.ReadFile("static/editor_start.html")
var footer, _ = ioutil.ReadFile("static/editor_end.html")

type EditorData struct {
	MarkdownIn  string
	MarkdownOut string
}

func editor(w http.ResponseWriter, r *http.Request) {
	markdown := r.FormValue("markdown")
	fmt.Println("in: " + markdown)
	unsafe := blackfriday.MarkdownCommon([]byte(markdown))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	var tmpl = template.Must(template.ParseFiles(
		templatesPath + "editor.html"))
	editorData := EditorData{markdown, string(html)}

	tmpl.ExecuteTemplate(w, "editor.html", editorData)
	//fmt.Fprint(w, string(header)+string(html)+string(footer))
}
