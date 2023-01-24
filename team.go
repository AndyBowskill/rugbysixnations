package main

import (
	"encoding/json"
	"net/http"
)

type Team struct {
	HeadCoach string   `json:"headcoach"`
	Forwards  []string `json:"forwards"`
	Backs     []string `json:"backs"`
}

func GetEnglandTeam(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/json")

	team := &Team{
		"Steve Borthwick",
		[]string{
			"Ollie Chessum", "Dan Cole", "Ben Curry", "Alex Dombrandt", "Ben Earl", "Ellis Genge", "Jamie George", "Joe Heyes", "Jonny Hill", "Nick Isiekwe", "Maro Itoje", "Courtney Lawes", "Lewis Ludlam", "George McGuigan", "Bevan Rodd", "Sam Simmonds", "Kyle Sinckler", "Mako Vunipola", "Jack Walker", "Jack Willis",
		},
		[]string{
			"Elliot Daly", "Owen Farrell", "Tommy Freeman", "Ollie Hassell-Collins", "Dan Kelly", "Max Malins", "Joe Marchant", "Alex Mitchell", "Cadan Murley", "Henry Slade", "Fin Smith", "Marcus Smith", "Freddie Steward", "Manu Tuilagi", "Jack van Poortvliet", "Ben Youngs",
		},
	}

	err := json.NewEncoder(w).Encode(team)
	if err != nil {
		panic(err)
	}
}
