package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/england", GetEnglandSquad)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

type Team struct {
	Player []Player `json:"player"`
}

type Player struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

func GetEnglandSquad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/json")

	Team := []Player{
		{Name: "Ellis Genge",
			Position: "Forward"},
		{Name: "Kyle Sinkler",
			Position: "Forward"},
		{Name: "Owen Farrel",
			Position: "Back"},
	}

	json.NewEncoder(w).Encode(Team)

}
