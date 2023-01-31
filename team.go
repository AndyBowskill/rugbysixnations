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

func jsonEncoder(resWri http.ResponseWriter, v any) {
	err := json.NewEncoder(resWri).Encode(v)
	if err != nil {
		panic(err)
	}
}

func getTeam(resWri http.ResponseWriter, req *http.Request, getTeamFromDB database) {

	setHeader(resWri)

	if req.Method == http.MethodGet {
		team := getTeamFromDB()
		jsonEncoder(resWri, team)

	} else {
		http.Error(resWri, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getEnglandTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() bson.M {
		return getTeamFromMongoDB("England")
	}

	getTeam(resWri, req, database)
}

func getScotlandTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() bson.M {
		return getTeamFromMongoDB("Scotland")
	}

	getTeam(resWri, req, database)
}

func getWalesTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() bson.M {
		return getTeamFromMongoDB("Wales")
	}

	getTeam(resWri, req, database)
}

func getIrelandTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() bson.M {
		return getTeamFromMongoDB("Ireland")
	}

	getTeam(resWri, req, database)
}
