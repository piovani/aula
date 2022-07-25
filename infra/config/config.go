package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	StageApp string `envconfig:"STAGE_APP"`

	ApiPort int `envconfig:"API_PORT"`
}

var Env Config

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
