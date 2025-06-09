package main

import (
	"flag"
	"log"
	"theztd/watchdog/internal/config"
	"theztd/watchdog/internal/logger"
	"theztd/watchdog/internal/probes"
	"theztd/watchdog/internal/server"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "Path to YAML config file")
	flag.Parse()

	logs := logger.InitLogger("DEBUG")
	logs.Info("Startuju aplikaci")

	// Example output
	cfg, err := config.Init(cfgPath)
	if err != nil {
		log.Println(err)
	}

	sharedStatus := probes.InitStateStorrage(cfg.Rules)

	go probes.RunCheckingAgent(cfg, sharedStatus)
	server.Run(sharedStatus)

}
