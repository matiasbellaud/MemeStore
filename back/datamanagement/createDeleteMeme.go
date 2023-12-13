package datamanagement

import (
	_ "github.com/lib/pq"
)

func CreateMeme(idUser int, memeName string, urlMeme string, description string, topic string) {
	// créer le meme dans la table meme
	AddDeleteUpdateDB("INSERT INTO meme (iduser, memename, urlmeme, description, topic, favourite) VALUES ($1,$2,$3,$4,$5,$6);", idUser, memeName, urlMeme, description, topic, false)
}
