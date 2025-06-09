package server

import (
	"encoding/json"
	"log"
	"net/http"
	"theztd/watchdog/internal/probes"
)

func Run(data *probes.Status) {
	http.HandleFunc("/_healthz/live", func(w http.ResponseWriter, r *http.Request) {
		data := data.Filter("live")
		json.NewEncoder(w).Encode(data)

	})

	http.HandleFunc("/_healthz/ready", func(w http.ResponseWriter, r *http.Request) {
		data := data.Filter("ready")
		log.Println(data)
		json.NewEncoder(w).Encode(data)

	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("HTTP server error: %v\n", err)
	}
}
