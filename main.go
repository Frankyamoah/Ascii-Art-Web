package main

import (
	"asciiart-web/asciifiles"
	"fmt"
	"net/http"
	"text/template"
)

type Text struct {
	normaltext string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*.html"))
	http.HandleFunc("/save", save)

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)

}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, _ := template.ParseFiles("standard.txt")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println(r.Form["asciitext"])
	}

	// inputtext := r.FormValue("usersinput")
	// outputtext := r.FormValue("asciitext")

	// data := struct {
	// 	Inputt  string
	// 	Outputt string
	// }{
	// 	Inputt:  inputtext,
	// 	Outputt: outputtext,
	// }

	// tpl.ExecuteTemplate(w, "processor.html", data)

}

func save(w http.ResponseWriter, r *http.Request) {
	fontval := r.FormValue("usersinput")
	fmt.Println(fontval)

	textval := r.FormValue("asciitext")
	fmt.Println(textval)

	fmt.Print(asciifiles.Asciiart(textval, fontval))
	//texts := &Text{normal}

	// b, err := json.Marshal(texts)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return

	// }
	// f, err := os.Open("/standard.txt")
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return

	// }
	// f.Write(b)
	// f.Close()
}
