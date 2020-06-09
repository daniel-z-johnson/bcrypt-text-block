package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "It works")
}

func main() {
    fmt.Println("Starting bcrypt-text-block")

    r := mux.NewRouter()

    r.HandleFunc("/", Home)
    r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
    http.ListenAndServe(":1117", r)
}
