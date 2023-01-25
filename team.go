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

func createTeam(headCoach string, forwards, backs []string) Team {
	return Team{
		HeadCoach: headCoach,
		Forwards:  forwards,
		Backs:     backs,
	}

}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/json")
}

func jsonEncoder(w http.ResponseWriter, team *Team) {
	err := json.NewEncoder(w).Encode(team)
	if err != nil {
		panic(err)
	}
}

func getEnglandTeam(w http.ResponseWriter, r *http.Request) {

	setHeader(w)

	if r.Method == http.MethodGet {

		team := &Team{
			"Steve Borthwick",
			[]string{
				"Ollie Chessum", "Dan Cole", "Ben Curry", "Alex Dombrandt", "Ben Earl", "Ellis Genge", "Jamie George", "Joe Heyes", "Jonny Hill", "Nick Isiekwe", "Maro Itoje", "Courtney Lawes", "Lewis Ludlam", "George McGuigan", "Bevan Rodd", "Sam Simmonds", "Kyle Sinckler", "Mako Vunipola", "Jack Walker", "Jack Willis",
			},
			[]string{
				"Elliot Daly", "Owen Farrell", "Tommy Freeman", "Ollie Hassell-Collins", "Dan Kelly", "Max Malins", "Joe Marchant", "Alex Mitchell", "Cadan Murley", "Henry Slade", "Fin Smith", "Marcus Smith", "Freddie Steward", "Manu Tuilagi", "Jack van Poortvliet", "Ben Youngs",
			},
		}
		jsonEncoder(w, team)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getScotlandTeam(w http.ResponseWriter, r *http.Request) {

	setHeader(w)

	if r.Method == http.MethodGet {

		team := &Team{
			"Gregor Townsend",
			[]string{
				"Ewan Ashman", "Josh Bayliss", "Simon Berghan", "Jamie Bhatti", "Fraser Brown", "Dave Cherry", "Andy Christie", "Luke Crosbie", "Jack Dempsey", "Matt Fagerson", "Zander Fagerson", "Grant Gilchrist", "Jonny Gray", "Richie Gray", "Cameron Henderson", "WP Nel", "Jamie Ritchie", "Pierre Schoeman", "Javan Sebastian", "Sam Skinner", "Rory Sutherland", "George Turner", "Hamish Watson",
			},
			[]string{
				"Chris Harris", "Ben Healy", "Stuart Hogg", "George Horne", "Huw Jones", "Blair Kinghorn", "Sean Maitland", "Ruaridh McConnochie", "Stafford McDowell", "Ali Price", "Cameron Redpath", "Finn Russell", "Ollie Smith", "Kyle Steyn", "Sione Tuipulotu", "Duhan van der Merwe", "Ben White",
			},
		}
		jsonEncoder(w, team)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getWalesTeam(w http.ResponseWriter, r *http.Request) {

	setHeader(w)

	if r.Method == http.MethodGet {

		team := &Team{
			"Warren Gatland",
			[]string{
				"Rhys Carre", "Wyn Jones", "Gareth Thomas", "Dewi Lake", "Ken Owens", "Bradley Roberts", "Leon Brown", "Tomas Francis", "Dillon Lewis", "Adam Beard", "Rhys Davies", "Dafydd Jenkins", "Alun Wyn Jones", "Teddy Williams", "Taulupe Faletau", "Jac Morgan", "Tommy Reffell", "Justin Tipuric", "Christ Tshiunza", "Aaron Wainwright",
			},
			[]string{
				"Kieran Hardy", "Rhys Webb", "Tomos Williams", "Dan Biggar", "Rhys Patchell", "Owen Williams", "Mason Grady", "Joe Hawkins", "George North", "Nick Tompkins", "Keiran Williams", "Josh Adams", "Alex Cuthbert", "Rio Dyer", "Leigh Halfpenny", "Louis Rees-Zammit", "Liam Williams",
			},
		}
		jsonEncoder(w, team)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getIrelandTeam(w http.ResponseWriter, r *http.Request) {

	setHeader(w)

	if r.Method == http.MethodGet {

		team := &Team{
			"Andy Farrell",
			[]string{
				"Aki", "R Byrne", "Casey", "Crowley", "Earls", "Gibson-Park", "Keenan", "Larmour", "Lowe", "McCloskey", "Murray", "O'Brien", "Osborne", "Ringrose", "Sexton", "Stockdale",
			},
			[]string{
				"Baird", "Bealham", "Beirne", "Conan", "Coombes", "Doris", "Furlong", "Healy", "Henderson", "Herring", "Kelleher", "Kilcoyne", "McCarthy", "O'Mahony", "O'Toole", "Porter", "Prendergast", "Ryan", "D Sheehan", "Van der Flier",
			},
		}
		jsonEncoder(w, team)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
