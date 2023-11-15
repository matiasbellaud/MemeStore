package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/connexion.html"))

	err := t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
