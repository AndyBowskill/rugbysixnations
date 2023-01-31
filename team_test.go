package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEnglandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/englandteam", nil)
	resRec := httptest.NewRecorder()

	englandHeadCoach := "Steve Borthwick"

	database := func() teamDatabase {
		var team teamDatabase
		team.HeadCoach = englandHeadCoach
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var team teamDatabase
	json.Unmarshal(bytes, &team)

	if team.HeadCoach != englandHeadCoach {
		t.Errorf("expected England head coach to be %s", englandHeadCoach)
	}
}

func TestPostEnglandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/englandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() teamDatabase {
		var team teamDatabase
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status == http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestGetScotlandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/scotlandteam", nil)
	resRec := httptest.NewRecorder()

	scotlandHeadCoach := "Gregor Townsend"

	database := func() teamDatabase {
		var team teamDatabase
		team.HeadCoach = scotlandHeadCoach
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var team teamDatabase
	json.Unmarshal(bytes, &team)

	if team.HeadCoach != scotlandHeadCoach {
		t.Errorf("expected Scotland head coach to be %s", scotlandHeadCoach)
	}
}

func TestPostScotlandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/scotlandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() teamDatabase {
		var team teamDatabase
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status == http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestGetWalesTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/walesteam", nil)
	resRec := httptest.NewRecorder()

	walesHeadCoach := "Warren Gatland"

	database := func() teamDatabase {
		var team teamDatabase
		team.HeadCoach = walesHeadCoach
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var team teamDatabase
	json.Unmarshal(bytes, &team)

	if team.HeadCoach != walesHeadCoach {
		t.Errorf("expected Wales head coach to be %s", walesHeadCoach)
	}
}

func TestPostWalesTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/walesteam", nil)
	resRec := httptest.NewRecorder()

	database := func() teamDatabase {
		var team teamDatabase
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status == http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestGetIrelandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/irelandteam", nil)
	resRec := httptest.NewRecorder()

	irelandHeadCoach := "Andy Farrell"

	database := func() teamDatabase {
		var team teamDatabase
		team.HeadCoach = irelandHeadCoach
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var team teamDatabase
	json.Unmarshal(bytes, &team)

	if team.HeadCoach != irelandHeadCoach {
		t.Errorf("expected Ireland head coach to be %s", irelandHeadCoach)
	}
}

func TestPostIrelandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/irelandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() teamDatabase {
		var team teamDatabase
		return team
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status == http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusMethodNotAllowed)
	}
}
