package main

import (
    "net/http" //package for http based web programs
    "fmt"
    "github.com/gorilla/mux"
)

func remindersHandler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Inside handler")
    fmt.Fprintf(w, "Hello world from my Go program!")
}

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/", remindersHandler)
    r.HandleFunc("/products", remindersHandler)
    r.HandleFunc("/articles", remindersHandler)
    http.Handle("/", r)
    http.ListenAndServe(":9000", nil)
}