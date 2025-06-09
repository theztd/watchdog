package config

type Config struct {
	Port          int `yaml:"port"`
	InitialStatus struct {
		Live  bool `yaml:"live"`
		Ready bool `yaml:"ready"`
	} `yaml:"initialStatus"`
	CheckIntervalSec int    `yaml:"checkIntervalSec"`
	Rules            []Rule `yaml:"rules"`
}

type Rule struct {
	Name     string   `yaml:"name"`
	ErrorMsg string   `yaml:"errorMsg"` // msg nebo name
	Address  string   `yaml:"address"`
	Port     int      `yaml:"port"`
	Timeout  int      `yaml:"timeout"`
	Method   string   `yaml:"method"`
	Required []string `yaml:"required"`
}
