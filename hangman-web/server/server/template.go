package server

import (
	"html/template"
	"net/http"
)

type Page struct {
	PageTitle string
	Mot       string
}

func Server(word string) {
	tmpl := template.Must(template.ParseFiles("test.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Page{
			PageTitle: "Test html",
			Mot:       word,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

/*go.mod: module server

go 1.21.0*/
