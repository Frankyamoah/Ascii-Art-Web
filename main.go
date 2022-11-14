package main

import (
	"asciiart-web/asciifiles"
	"net/http"
	"strings"
	"text/template"
)

type Text struct {
	Normaltext string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*.html"))
	// http.HandleFunc("/save", save)

}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/process", Processor)
	http.ListenAndServe(":8080", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)

}

func Processor(w http.ResponseWriter, r *http.Request) {

	fontval := r.FormValue("usersinput")
	splitfont := strings.Split(fontval, "")

	textval := r.FormValue("asciitext")
	splittxt := strings.Split(textval, "")

	text := asciifiles.Asciiart(splittxt, splitfont)

	data := Text{
		Normaltext: text,
	}
	tpl.ExecuteTemplate(w, "index.html", data)
	//fmt.Print(asciifiles.Asciiart(textval, fontval))

}
