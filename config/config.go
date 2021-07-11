package config

import (
	"flag"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	API struct {
		URL string `yaml:"url"`
		UserID string `yaml:"userid"`
		Password string `yaml:"password"`
	}
}

type Flags struct {
	ConfigFile string // Path to yaml config file
}

func ParseFlags() *Flags {
	f := new(Flags)
	flag.StringVar(
		&f.ConfigFile,
		"config",
		"vmware_exporter.yml",
		"Path to the vmware_exporter configuration file",
	)
	flag.Parse()
	return f
}

// Create a global variable to hold the configuration
var cfg *Config

func ParseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	y := yaml.NewDecoder(file)
	config := new(Config)
	if err := y.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
