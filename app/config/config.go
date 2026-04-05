package config

import (
	"encoding/json"
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

type Config struct {
	port      string
	packSizes []int
}

func LoadConfig() Config {
	cfg := Config{
		port:      DefaultPort,
		packSizes: DefaultPackSizes,
	}

	cfgFile, err := os.Open("config.json")
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

	if temp.Port != "" {
		cfg.port = temp.Port
	}

	if len(temp.Packs) > 0 {
		slices.Sort(temp.Packs)
		cfg.packSizes = temp.Packs
	}

	return cfg
}

func (c Config) Port() string {
	return c.port
}

func (c Config) PackSizes() []int {
	return c.packSizes
}
