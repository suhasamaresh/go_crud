package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Nbaplayer struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Team     string `json:"team"`
	Coach    *coach `json:"coach"`
}

type coach struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var nbaplayers []Nbaplayer

func getnbaplayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(nbaplayers)
}

func deletenbaplayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range nbaplayers {

		if item.Name == params["name"] {
			nbaplayers = append(nbaplayers[:index], nbaplayers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(nbaplayers)
}

func getnbaplayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range nbaplayers {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createnbaplayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var nbaplayer Nbaplayer
	_ = json.NewDecoder(r.Body).Decode(&nbaplayer)
	nbaplayers = append(nbaplayers, nbaplayer)
	json.NewEncoder(w).Encode(nbaplayer)
}

func updatenbaplayer(w http.ResponseWriter, r *http.Request) {
	//set content type
	w.Header().Set("Content-type", "application-json")
	//params
	params := mux.Vars(r)
	//loop over nbaplayers, range
	//delete the nbaplayer with the name you have sent
	//add a new nbaplayer-the movie that we send in postman
	for index, item := range nbaplayers {
		if item.Name == params["name"] {
			nbaplayers = append(nbaplayers[:index], nbaplayers[index+1:]...)
			var nbaplayer Nbaplayer
			_ = json.NewDecoder(r.Body).Decode(&nbaplayer)
			nbaplayer.Name = params["name"]
			nbaplayers = append(nbaplayers, nbaplayer)
			json.NewEncoder(w).Encode(nbaplayer)
		}
	}

}

func main() {

	r := mux.NewRouter()

	nbaplayers = append(nbaplayers, Nbaplayer{Name: "Michael jordan", Position: "Shooting guard", Team: "Chicago bulls", Coach: &coach{Firstname: "Phil", Lastname: "Jackson"}})
	nbaplayers = append(nbaplayers, Nbaplayer{Name: "Stephen Curry", Position: "Point guard", Team: "Golden State Warriors", Coach: &coach{Firstname: "Steve", Lastname: "Kerr"}})
	nbaplayers = append(nbaplayers, Nbaplayer{Name: "Shaq ", Position: "Center", Team: "Los Angeles Lakers", Coach: &coach{Firstname: "Phil", Lastname: "Jackson"}})
	r.HandleFunc("/Nbaplayer", getnbaplayers).Methods("GET")
	r.HandleFunc("/Nbaplayer/{name}", getnbaplayer).Methods("GET")
	r.HandleFunc("/Nbaplayer", createnbaplayer).Methods("POST")
	r.HandleFunc("/Nbaplayer/{name}", updatenbaplayer).Methods("PUT")
	r.HandleFunc("/Nbapleyer/{name}", deletenbaplayer).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
