package main

import (
    "net/http" //package for http based web programs
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Reminder struct {
	Date string
	Type string
}

func (r Reminder) String() string {
	return r.Date + " " + r.Type
}

func remindersGetHandler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Inside handler")
    fmt.Fprintf(w, "Hello world from my Go program!")
}

func remindersPostHandler(w http.ResponseWriter, r *http.Request) { 
    decoder := json.NewDecoder(r.Body)
    var t Reminder
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }   
    fmt.Println(t)
    fmt.Fprintf(w, "Hello world from my Go program!")
}


func main() {
	r := mux.NewRouter()
    r.HandleFunc("/reminders", remindersGetHandler).
    	Methods("GET")
     r.HandleFunc("/reminders", remindersPostHandler).
    	Methods("POST")

    http.Handle("/", r)
    http.ListenAndServe(":9000", nil)
}