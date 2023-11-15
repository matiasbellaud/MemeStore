package main

import (
	"fmt"
	"memestore/back/handler"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

func main() {
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("front/css"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("front/assets"))))
	fmt.Println("(http://localhost"+port+"/"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
