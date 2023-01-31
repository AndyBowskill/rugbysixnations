package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateTeam(t *testing.T) {

	headCoach := "Steve Borthwick"
	forwards := []string{"Ellis Genge"}
	backs := []string{"Owen Farrell"}

	expected := Team{
		HeadCoach: headCoach,
		Forwards:  forwards,
		Backs:     backs,
	}

	result := createTeam(headCoach, forwards, backs)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestGetEnglandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/englandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() bson.M {
		return nil
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestPostEnglandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/englandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() bson.M {
		return nil
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

	database := func() bson.M {
		return nil
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestPostScotlandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/scotlandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() bson.M {
		return nil
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

	database := func() bson.M {
		return nil
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestPostWalesTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/walesteam", nil)
	resRec := httptest.NewRecorder()

	database := func() bson.M {
		return nil
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

	database := func() bson.M {
		return nil
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusOK)
	}

	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestPostIrelandTeam(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/irelandteam", nil)
	resRec := httptest.NewRecorder()

	database := func() bson.M {
		return nil
	}

	getTeam(resRec, req, database)

	res := resRec.Result()
	defer res.Body.Close()

	if status := resRec.Code; status == http.StatusOK {
		t.Errorf("expected status code to be %v, want %v",
			status, http.StatusMethodNotAllowed)
	}
}
