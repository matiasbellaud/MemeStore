package main

import (
    "memestore/back/handler"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {

    file, err := os.OpenFile("acount.txt", os.O_WRONLY|os.O_APPEND, 0600)
    defer file.Close() // on ferme automatiquement à la fin de notre programme

    if err != nil {
        panic(err)
    }

    hashd, err := handler.HasingFunc("password")
    if err != nil {
        fmt.Println("1",err)
        return
    }

    _, err = file.WriteString(hashd+"\n")
    if err != nil {
        panic(err)
    }

    data, err := ioutil.ReadFile("acount.txt") // lire le fichier
    if err != nil {
        fmt.Println(err)
    }

    boolean, err := handler.PasswordEqualHash("password3", string(data))
    if err != nil {
        fmt.Println("2",err)
        return
    }

    fmt.Print(string(data))
    fmt.Println(boolean)
    return
}