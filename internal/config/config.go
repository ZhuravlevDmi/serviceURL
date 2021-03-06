package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

var ListContentType = []string{"application/javascript", "application/json",
	"text/css", "text/html", "text/plain", "text/xml"}

type ConfigAdress struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	PATHFile      string `env:"FILE_STORAGE_PATH"`
}

//envDefault:"file.txt" envExpand:"true"
func (c *ConfigAdress) Parse() {
	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "ServerAddress")
	flag.StringVar(&c.BaseURL, "b", c.BaseURL, "BaseURL")
	flag.StringVar(&c.PATHFile, "f", c.PATHFile, "PATHFile")
	flag.Parse()
}

func (c *ConfigAdress) ParseTest() {
	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}
}
