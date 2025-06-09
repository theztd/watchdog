package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Init(cfgPath string) (cfg Config, error error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", cfgPath, err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing YAML: %v\n", err)
		os.Exit(1)
	}

	return cfg, nil
}
