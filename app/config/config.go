package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
)

const DefaultPort = "8080"

var (
	DefaultPackSizes = []int{
		250, 500, 1000, 2000, 5000,
	}
)

// Config represents necessary configurations for the application, namely port and available pack sizes.
type Config struct {
	port      string
	packSizes []int
}

// LoadConfig creates an instance of Config and populates it whether from `config.json` file or from Defaults.
func LoadConfig(filename string) Config {
	if filename == "" {
		filename = "config.json"
	}

	cfg := Config{
		port:      DefaultPort,
		packSizes: DefaultPackSizes,
	}

	cfgFile, err := os.Open(filename)
	if err != nil {
		log.Printf("Error while reading config file: %s\n", err)
	}
	defer cfgFile.Close()

	temp := struct {
		Port  string `json:"port"`
		Packs []int  `json:"pack_sizes"`
	}{}

	decoder := json.NewDecoder(cfgFile)
	decoder.Decode(&temp)

	fmt.Println(temp)

	if temp.Port != "" {
		cfg.port = temp.Port
	}

	if len(temp.Packs) > 0 {
		slices.Sort(temp.Packs)
		cfg.packSizes = temp.Packs
	}

	return cfg
}

// Port returns string representation of port number.
func (c Config) Port() string {
	return c.port
}

// PackSizes returns slice of available pack sizes sorted increasingly.
func (c Config) PackSizes() []int {
	return c.packSizes
}
