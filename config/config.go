package config

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

const defaultName = "./config/app.toml"

type Config struct {
	Name     string    `toml:"name"`
	HTTP     *HTTP     `toml:"http"`
	Database *Database `toml:"database"`
	Env      Env       `toml:"env"`
}

type HTTP struct {
	Port string `toml:"port"`
}

type Database struct {
	System string `toml:"system"`
}

type Env map[string]Variable

func (e Env) Get(name string) string {
	if v, ok := e[name]; ok {
		return v.Value
	}

	return ""
}

type Variable struct {
	Value string `toml:"value"`
}

func Parse() (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile(defaultName, &config); err != nil {
		return nil, err
	}

	for key, _ := range config.Env {
		upperKey := strings.ToUpper(key)
		if value, exists := os.LookupEnv(upperKey); exists {
			config.Env[key] = Variable{Value: value}
		}
	}

	return &config, nil
}
