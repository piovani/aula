package main

import (
	"log"

	"github.com/piovani/aula/api"
	"github.com/piovani/aula/infra/config"
)

func main() {
	var err error

	err = config.StarConfig()
	FatalError(err)

	err = api.NewService().Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
