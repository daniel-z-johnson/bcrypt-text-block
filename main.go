package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "golang.org/x/crypto/bcrypt"
)

var minuteCost int

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "It works")
}

func favicon(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/static/favicon.ico", http.StatusFound)
}

func main() {
    fmt.Println("Starting bcrypt-text-block")
    findMinuteCost()
    fmt.Println(minuteCost)
    r := mux.NewRouter()

    r.HandleFunc("/", Home)
    r.HandleFunc("/favicon.ico", favicon)
    r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
    http.ListenAndServe(":1117", r)
}

func findMinuteCost() {
   start := time.Now()
   bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
   dur := time.Since(start)
   cost := bcrypt.DefaultCost
   for dur < time.Minute {
       cost += 1
       dur += dur
   }
   minuteCost = cost
}


