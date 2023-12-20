package datamanagement

import (
	"log"

	_ "github.com/lib/pq"
)

func RecuperationIdUser(mail string) int {
	var IdUser int
	rows := SelectDB("SELECT idUser FROM account  WHERE Mail = $1;", mail)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&IdUser)
		if err != nil {
			log.Fatal(err)
		}
	}

	return IdUser
}

// SELECT COUNT(*) FROM table_name WHERE condition;
func CountUser(mail string) int {
	var count int
	rows := SelectDB("SELECT COUNT(*) FROM account WHERE mail = $1;", mail)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}

	return count
}

func RecuperationInfoUser(mail string) []UserFromDB {
	var User []UserFromDB
	var iduser int
	var description string
	var username string
	var password string
	var verificationword string

	//recupère un utilisateur de la table account
	rows := SelectDB("SELECT iduser , description  FROM account  WHERE mail = $1;", mail)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&iduser, &description)
		if err != nil {
			log.Fatal(err)
		}
	}

	//recupère la table infosecuraccount
	rows = SelectDB("SELECT username, password, verificationword FROM infosecuraccount  WHERE iduser = $1;", iduser)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&username, &password, &verificationword)
		if err != nil {
			log.Fatal(err)
		}
	}

	UserStruct := UserFromDB{
		IdUser:           iduser,
		Description:      description,
		Username:         username,
		Password:         password,
		VerificationWord: verificationword,
	}
	User = append(User, UserStruct)
	return User
}
