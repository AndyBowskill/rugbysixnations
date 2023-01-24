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

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
