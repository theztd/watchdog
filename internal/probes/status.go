package probes

import (
	"fmt"
	"sync"
	"theztd/watchdog/internal/config"
)

type Result struct {
	Counter    int
	LastStatus string
	Message    string
	Required   []string
}

/*
Struktura drzici stav prob s mutex zamkem RW.
*/
type Status struct {
	Results map[string]Result

	RW sync.Mutex
}

func InitStateStorrage(rules []config.Rule) *Status {
	results := map[string]Result{}

	for _, rule := range rules {
		results[rule.Name] = Result{Counter: 0, LastStatus: "init", Required: rule.Required}
	}
	return &Status{
		// musim vytvorit map[jobs]Result, kde vse je down.
		Results: results,
	}
}

func (s *Status) Update(rule config.Rule, status string) {
	s.RW.Lock()
	defer s.RW.Unlock()

	last := s.Results[rule.Name]
	if last.LastStatus == status {
		last.Counter++
	} else {
		last.LastStatus = status
		last.Counter = 1
	}
	s.Results[rule.Name] = last
}

func (s *Status) GetAll() map[string]Result {
	s.RW.Lock()
	defer s.RW.Unlock()
	return s.Results
}

func contains(word string, slice []string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}

func (s *Status) Filter(byProbe string) (results map[string]Result) {
	s.RW.Lock()
	defer s.RW.Unlock()

	for name, r := range s.Results {
		if contains(byProbe, r.Required) {
			fmt.Println("Switch: contains ", r)
			results[name] = r
		} else {
			fmt.Println("Switch: NO contains ", r)
		}
	}

	return results
}
