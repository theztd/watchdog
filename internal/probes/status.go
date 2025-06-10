package probes

import (
	"sync"
	"theztd/watchdog/internal/config"
)

type Result struct {
	Counter    int      `json:"-"`
	LastStatus string   `json:"Status"`
	Required   []string `json:"-"`
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

// Pomocna funkce
func contains(word string, slice []string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}

func (s *Status) Filter(byProbe string) map[string]Result {
	s.RW.Lock()
	defer s.RW.Unlock()

	results := make(map[string]Result)

	for name, data := range s.Results {
		if contains(byProbe, data.Required) {
			results[name] = data
		}
	}

	return results
}
