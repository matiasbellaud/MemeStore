package handler

import (
	"html/template"
	"log"
	"memestore/back/datamanagement"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func Inscription(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/inscription.html"))

	email := r.FormValue("userEmail")
	username := r.FormValue("username")
	userPassword := r.FormValue("userPassword")
	userVerificationWord := r.FormValue("userVerificationWord")

	if email != "" && username != "" && userPassword != "" && userVerificationWord != "" {

		goodCreation, userConnected := datamanagement.CreateAccount("", email, username, userPassword, userVerificationWord)
		if goodCreation {
			//creation de cookie pour le reste du site
			cookieIdUser := http.Cookie{Name: "idUser", Value: strconv.Itoa(userConnected[0].IdUser), Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIdUser)

			cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true", Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIsConnected)

			http.Redirect(w, r, "/Sites/", http.StatusSeeOther)
		}
	}

	err := t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
