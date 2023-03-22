package main

import (
	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"Director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Movies", getmovie).Methods("GET")
	r.HandleFunc("/Movies/[id]", gertmovie).Methods("GET")
	r.HandleFunc("/Movies", createmovie).Methods("POST")
	r.HandleFunc("/Movies/[id]", updatemovie).Methods("PUT")
	r.HandleFunc("Movie[id]", deletemovie).Methods("DELETE")
}
