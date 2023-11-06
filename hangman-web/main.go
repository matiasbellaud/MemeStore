package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)

	fmt.Println("(http://localhost:8080) - Started on port, port")
	http.ListenAndServe(port, nil)
}
