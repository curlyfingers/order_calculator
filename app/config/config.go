package config

import "slices"

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
	return Config{}
}

func (c Config) Port() string {
	if c.port != "" {
		return c.port
	}

	return DefaultPort
}

func (c Config) PackSizes() []int {
	if len(c.packSizes) > 0 {
		slices.Sort(c.packSizes)
		return c.packSizes
	}

	return DefaultPackSizes
}
