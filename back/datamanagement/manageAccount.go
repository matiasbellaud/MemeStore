package datamanagement

import (
	_ "github.com/lib/pq"
)

func CreateAccount(description string, mail string, username string, password string, verificationWord string) {
	// créer le compte dans la table account
	AddDeleteUpdateDB("INSERT INTO account (description, mail) VALUES ($1,$2);", description, mail)
	// viens chercher l'id du compte créer au dessus
	idUser := RecuperationIdUser(mail)
	// créer le compte dans la table accountSecurInfo
	AddDeleteUpdateDB("INSERT INTO InfoSecurAccount (IdUser, Username, Password, VerificationWord) VALUES ($1,$2,$3,$4);", idUser, username, password, verificationWord)
}
