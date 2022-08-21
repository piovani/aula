package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiPort int `envconfig:"API_PORT"`

	MongURL      string `envconfig:"MONGO_URL"`
	MongDatabase string `envconfig:"MONGO_DATABASE"`
}

var Env Config

func StarConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
