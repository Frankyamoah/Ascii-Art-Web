package main

import (
	"asciiart-web/asciifiles"
	"fmt"
	"net/http"
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
	/*
		if r.Method == "POST" {
			t, _ := template.ParseFiles("index.html")
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			fmt.Println(r.Form["asciitext"])
		}*/

	fontval := r.FormValue("usersinput")
	//fmt.Println(fontval)

	textval := r.FormValue("asciitext")
	//fmt.Println(textval)
	text := asciifiles.Asciiart(textval, fontval)

	data := Text{
		Normaltext: text,
	}
	tpl.ExecuteTemplate(w, "index.html", data)
	fmt.Print(asciifiles.Asciiart(textval, fontval))

}

// func save(w http.ResponseWriter, r *http.Request) {
// 	fontval := r.FormValue("usersinput")
// 	//fmt.Println(fontval)

// 	textval := r.FormValue("asciitext")
// 	//fmt.Println(textval)

// 	data := Text{
// 		normaltext: asciifiles.Asciiart(textval, fontval),
// 	}
// 	tpl.ExecuteTemplate(w, "index.html", data)
// 	fmt.Print(asciifiles.Asciiart(textval, fontval))

// }
