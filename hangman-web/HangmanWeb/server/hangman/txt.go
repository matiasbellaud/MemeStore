package server

import (
	"math/rand"
	"os"
)

func Start() ([]string, []string ) {
	soluce := SelectWord()
	var index int
	word := []string{}
	for _, i := range soluce {
		word = append(word, i)
	}
	for i := 0; i < len(word); i++ {
		index = rand.Intn(len(word))
		for word[index] == "_" {
			index = rand.Intn(len(word))
		}
		word[index] = "_"
	}
	return soluce, word
}


func SelectWord() []string {
	contenu, _ := os.ReadFile("words.txt")
	words := byteToString(contenu)
	mot := words[rand.Intn(len(words))]
	run := []rune(mot)
	soluce := []string{}
	for i := 0; i < len(run)-1; i++ {
		soluce = append(soluce, string(run[i]))
	}
	return soluce
}

func byteToString(b []byte) []string {
	word := ""
	result := []string{}
	for _, i := range b {
		if string(i) == "\n" {
			result = append(result, word)
			word = ""
		} else {
			word += string(i)
		}
	}
	return result
}

func ListToString(s []string) string {
	word := ""
	for _, i := range s {
		word += i
	}
	return word
}