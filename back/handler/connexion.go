package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/connexion.html"))

	// connection a la BDD
	db, err := sql.Open("postgres", "user=bella password=171104 host=localhost dbname=test_pgadmin sslmode=disable")
	if err != nil {
		log.Fatalf("Error : Unable to connect to database : %v", err)
	}

	rows, err := db.Query("SELECT * FROM Account")
	if err != nil {
		log.Fatalf("Error : Unable to execute query : %v", err)
	} else {
		print("maybe")
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("User ID : %d, Name : %s\n", id, name)
	}

	err = t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
