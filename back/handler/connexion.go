package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/connexion.html"))

	//datamanagement.CreateAccount("test de la description", "test@test", "belanus", "pass", "verif")

	//mettre en place les cookie pour avoir des variables globals
	cookieIdUser := http.Cookie{Name: "idUser", Value: strconv.Itoa(2), Expires: time.Now().Add(30 * time.Minute)}
	http.SetCookie(w, &cookieIdUser)

	err := t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
