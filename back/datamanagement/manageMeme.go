package datamanagement

import (
	"log"

	_ "github.com/lib/pq"
)

func UpdateLike(idMeme int, isLike bool) {
	//changer la valeur de favorite dans la BDD
	AddDeleteUpdateDB("UPDATE meme SET favorite = $1 WHERE idmeme = $2;", isLike, idMeme)
}

func RecuperationMeme(iduser int) []MemeFromDB {
	var allMemeUser []MemeFromDB

	//recupère les meme de la table account
	rows := SelectDB("SELECT idmeme ,iduser, memename,urlmeme,description, topic, favorite  FROM meme  WHERE iduser = $1;", iduser)
	defer rows.Close()

	for rows.Next() {
		var idmeme int
		var idusermeme int
		var memename string
		var urlmeme string
		var description string
		var topic string
		var favorite int
		err := rows.Scan(&idmeme, &idusermeme, &memename, &urlmeme, &description, &topic, &favorite)
		if err != nil {
			log.Fatal(err)
		}
		memeUserStruct := MemeFromDB{
			IdMeme:      idmeme,
			IdUser:      idusermeme,
			MemeName:    memename,
			UrlMeme:     urlmeme,
			Description: description,
			Topic:       topic,
			Favorite:    favorite,
		}
		allMemeUser = append(allMemeUser, memeUserStruct)
	}

	return allMemeUser
}
