package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var minuteCost int

type BcryptTextBlock struct {
	Bcrypt  string `json:"bcrypt"`
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It works")
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon.ico", http.StatusFound)
}

func BcryptEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		body := BcryptTextBlock{Message: err.Error()}
		b, err := json.Marshal(body)
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			return
		}
		w.Write(b)
		return
	}
	// should only be one message block should have more error
	block := r.Form["textBlock"][0]
	fmt.Printf("block of text: %s\n", block)
	bcryptText, err := bcrypt.GenerateFromPassword([]byte(block), minuteCost)
	fmt.Printf("bcrypt: %s\n", bcryptText)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := BcryptTextBlock{Bcrypt: string(bcryptText)}
	b, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func main() {
	fmt.Println("Starting bcrypt-text-block")
	findMinuteCost()
	fmt.Println(minuteCost)
	r := mux.NewRouter()

	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/favicon.ico", favicon).Methods("GET")
	r.HandleFunc("/bcrypt-this", BcryptEndPoint).Methods("POST")
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
	http.ListenAndServe(":1117", r)
}

func findMinuteCost() {
	start := time.Now()
	bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	dur := time.Since(start)
	cost := bcrypt.DefaultCost
	for dur < time.Minute {
		cost++
		dur += dur
	}
	minuteCost = cost
}
