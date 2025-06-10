package probes

import (
	"fmt"
	"log"
	"sync"
	"theztd/watchdog/internal/config"
	"theztd/watchdog/internal/httpCheck"
	"theztd/watchdog/internal/ping"
	"time"
)

func RunCheckingAgent(cfg config.Config, status *Status) {
	type resultsChan struct {
		Rule   config.Rule
		Status error
	}
	for {
		var wg sync.WaitGroup
		resultsCh := make(chan resultsChan, len(cfg.Rules))

		// Spustime vsechny checky paralelne
		for _, rule := range cfg.Rules {
			wg.Add(1)
			go func(rule config.Rule) {
				defer wg.Done()
				ok := checkRule(rule)
				log.Println("DEBUG", ok, rule.Name)
				resultsCh <- resultsChan{Rule: rule, Status: ok}
			}(rule)
		}

		wg.Wait()
		close(resultsCh)

		// Vysledky vsech checku zpracovavam jak dorazi
		for res := range resultsCh {
			if res.Status != nil {
				status.Update(res.Rule, res.Status.Error())
			} else {
				status.Update(res.Rule, "Ok")
			}

		}

		time.Sleep(time.Duration(cfg.CheckIntervalSec) * time.Second)
	}
}

func checkRule(rule config.Rule) error {
	switch rule.Method {
	case "ping":
		if err := ping.CheckTCPPort(rule.Address, rule.Port, time.Duration(2)*time.Second); err != nil {
			log.Printf("ERR [ping]: Test failed (%s)", rule.ErrorMsg)
			return err
		} else {
			return nil
		}

	case "resolve":
		// TODO: later
		fmt.Println("___ TODO: Doing DNS resolve (need implementation in checks.go)")
		return nil

	case "http-get":
		_, err := httpCheck.GetV2(rule.Address)
		if err != nil {
			log.Printf("ERR [http-get]: Test failed (%s)", rule.ErrorMsg)
			return fmt.Errorf("ERR [http-get]: Test failed (%s)", rule.ErrorMsg)
		}
		// pokud mam zaple metrics mohu reportovat metriky odpovedi, staci pouzit response z GetV2
		return nil

	default:
		fmt.Println("Unknown method")
		return fmt.Errorf("probes.checkRule Unknown method ")
	}
}
