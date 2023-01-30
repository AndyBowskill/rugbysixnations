package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
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

func setHeader(resWri http.ResponseWriter) {
	resWri.Header().Set("Content-Type", "text/json")
}

func jsonEncoder(resWri http.ResponseWriter, team *Team) {
	err := json.NewEncoder(resWri).Encode(team)
	if err != nil {
		panic(err)
	}
}

func jsonEncoderTeam(resWri http.ResponseWriter, v any) {
	err := json.NewEncoder(resWri).Encode(v)
	if err != nil {
		panic(err)
	}
}

func englandTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() bson.M {
		return getTeamFromMongoDB("England")
	}

	getEnglandTeam(resWri, req, database)
}

func getEnglandTeam(resWri http.ResponseWriter, req *http.Request, englandTeam database) {

	setHeader(resWri)

	if req.Method == http.MethodGet {
		team := englandTeam()
		jsonEncoderTeam(resWri, team)

	} else {
		http.Error(resWri, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getScotlandTeam(resWri http.ResponseWriter, req *http.Request) {

	setHeader(resWri)

	if req.Method == http.MethodGet {

		team := &Team{
			"Gregor Townsend",
			[]string{
				"Ewan Ashman", "Josh Bayliss", "Simon Berghan", "Jamie Bhatti", "Fraser Brown", "Dave Cherry", "Andy Christie", "Luke Crosbie", "Jack Dempsey", "Matt Fagerson", "Zander Fagerson", "Grant Gilchrist", "Jonny Gray", "Richie Gray", "Cameron Henderson", "WP Nel", "Jamie Ritchie", "Pierre Schoeman", "Javan Sebastian", "Sam Skinner", "Rory Sutherland", "George Turner", "Hamish Watson",
			},
			[]string{
				"Chris Harris", "Ben Healy", "Stuart Hogg", "George Horne", "Huw Jones", "Blair Kinghorn", "Sean Maitland", "Ruaridh McConnochie", "Stafford McDowell", "Ali Price", "Cameron Redpath", "Finn Russell", "Ollie Smith", "Kyle Steyn", "Sione Tuipulotu", "Duhan van der Merwe", "Ben White",
			},
		}
		jsonEncoder(resWri, team)

	} else {
		http.Error(resWri, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getWalesTeam(resWri http.ResponseWriter, req *http.Request) {

	setHeader(resWri)

	if req.Method == http.MethodGet {

		team := &Team{
			"Warren Gatland",
			[]string{
				"Rhys Carre", "Wyn Jones", "Gareth Thomas", "Dewi Lake", "Ken Owens", "Bradley Roberts", "Leon Brown", "Tomas Francis", "Dillon Lewis", "Adam Beard", "Rhys Davies", "Dafydd Jenkins", "Alun Wyn Jones", "Teddy Williams", "Taulupe Faletau", "Jac Morgan", "Tommy Reffell", "Justin Tipuric", "Christ Tshiunza", "Aaron Wainwright",
			},
			[]string{
				"Kieran Hardy", "Rhys Webb", "Tomos Williams", "Dan Biggar", "Rhys Patchell", "Owen Williams", "Mason Grady", "Joe Hawkins", "George North", "Nick Tompkins", "Keiran Williams", "Josh Adams", "Alex Cuthbert", "Rio Dyer", "Leigh Halfpenny", "Louis Rees-Zammit", "Liam Williams",
			},
		}
		jsonEncoder(resWri, team)

	} else {
		http.Error(resWri, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getIrelandTeam(resWri http.ResponseWriter, req *http.Request) {

	setHeader(resWri)

	if req.Method == http.MethodGet {

		team := &Team{
			"Andy Farrell",
			[]string{
				"Aki", "R Byrne", "Casey", "Crowley", "Earls", "Gibson-Park", "Keenan", "Larmour", "Lowe", "McCloskey", "Murray", "O'Brien", "Osborne", "Ringrose", "Sexton", "Stockdale",
			},
			[]string{
				"Baird", "Bealham", "Beirne", "Conan", "Coombes", "Doris", "Furlong", "Healy", "Henderson", "Herring", "Kelleher", "Kilcoyne", "McCarthy", "O'Mahony", "O'Toole", "Porter", "Prendergast", "Ryan", "D Sheehan", "Van der Flier",
			},
		}
		jsonEncoder(resWri, team)

	} else {
		http.Error(resWri, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
