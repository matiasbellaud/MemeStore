package main

import (
	server "server/server"
	pendu "server/hangman"
	"fmt"
)

func main() {
	mot := pendu.ListToString(pendu.SelectWord())
	fmt.Println(mot)
	server.Server(mot)
}