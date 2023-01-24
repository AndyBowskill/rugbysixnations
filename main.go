package main

import (
	"net/http"
	"time"
)

func main() {

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 + time.Second,
		WriteTimeout: 30 + time.Second,
		IdleTimeout:  120 + time.Second,
	}

	http.HandleFunc("/englandteam", GetEnglandTeam)
	http.HandleFunc("/scotlandteam", GetScotlandTeam)
	http.HandleFunc("/walesteam", GetWalesTeam)
	http.HandleFunc("/irelandteam", GetIrelandTeam)

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
