package handler

import (
	"html/template"
	"log"
	"memestore/back/datamanagement"
	"net/http"

	_ "github.com/lib/pq"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/connexion.html"))

	datamanagement.CreateAccount("test de la description", "test@test", "belanus", "pass", "verif")

	err := t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
