package main

import (
	"encoding/json"
	"net/http"
)

func setHeader(resWri http.ResponseWriter) {
	resWri.Header().Set("Content-Type", "text/json")
}

func jsonEncoder(resWri http.ResponseWriter, team teamDatabase) {
	err := json.NewEncoder(resWri).Encode(team)
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

	database := func() teamDatabase {
		return getTeamFromMongoDB("England")
	}

	getTeam(resWri, req, database)
}

func getScotlandTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() teamDatabase {
		return getTeamFromMongoDB("Scotland")
	}

	getTeam(resWri, req, database)
}

func getWalesTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() teamDatabase {
		return getTeamFromMongoDB("Wales")
	}

	getTeam(resWri, req, database)
}

func getIrelandTeam(resWri http.ResponseWriter, req *http.Request) {

	database := func() teamDatabase {
		return getTeamFromMongoDB("Ireland")
	}

	getTeam(resWri, req, database)
}
