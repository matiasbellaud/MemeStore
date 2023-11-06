package server

import (
	"fmt"
	"os"
)

func Turn(word, soluce []string) ([]string, int, string) {
	var input string
	var cd int
	_, err := fmt.Scan(&input)
	if err != nil {
		os.Exit(0)
	}
	if len(input) != 1 {
		os.Exit(0)
	}
	if input == "+" {
		os.Exit(0)
	}
	l := check(input,soluce)
	if len(l) == 0 {
		cd = -1
	} else {
		cd = 0
	}
	for _, i := range l {
		word[i] = soluce[i]
	}
	return word, cd, input
	
}

func check(s string, soluce []string) []int {
	l := []int{}
	for f, i := range soluce {
		if s == i {
			l = append(l, f)
		}
	}
	return l
}

func Victory(word, soluce []string) bool {
	for f, i := range soluce {
		if word[f] != i {
			return false
		}
	}
	return true
}