package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	w := httptest.NewRecorder()

	getEnglandTeam(w, req)

	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
