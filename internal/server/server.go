package server

import (
	"encoding/json"
	"log"
	"net/http"
	"theztd/watchdog/internal/probes"
	"time"
)

func Run(data *probes.Status) {
	http.HandleFunc("/_healthz/live", func(w http.ResponseWriter, r *http.Request) {
		data := data.Filter("live")

		respStatus := 200
		for _, v := range data {
			if v.LastStatus != "Ok" {
				respStatus = 503
				break
			}
		}

		w.WriteHeader(respStatus)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("ERR [server.live]: ", err)
		}

	})

	http.HandleFunc("/_healthz/ready", func(w http.ResponseWriter, r *http.Request) {
		data := data.Filter("ready")

		var respStatus int = 200
		for _, v := range data {
			if v.LastStatus != "Ok" {
				respStatus = 503
				break
			}
		}

		w.WriteHeader(respStatus)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("ERR [server.live]: ", err)
		}

	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // nebo tvůj mux
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
		IdleTimeout:  5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v\n", err)
	}
}
