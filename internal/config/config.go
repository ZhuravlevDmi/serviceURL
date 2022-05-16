package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type ConfigAdress struct {
	User          string `env:"USER"`
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	PATHFile      string `env:"FILE_STORAGE_PATH"`
}

type FlagStruct struct {
	ServerAddress *string
	BaseURL       *string
	PATHFile      *string
}

//envDefault:"file.txt" envExpand:"true"
func (c *ConfigAdress) Parse() {
	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}
}
