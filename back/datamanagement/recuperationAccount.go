package datamanagement

import (
	"log"

	_ "github.com/lib/pq"
)

func RecuperationIdUser(mail string) int {
	var IdUser int
	rows := SelectDB("SELECT IdUser FROM Account  WHERE Mail = $1;", mail)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&IdUser)
		if err != nil {
			log.Fatal(err)
		}
	}

	return IdUser
}
