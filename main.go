package main

import (
	"log"

	"github.com/piovani/aula/api"
	"github.com/piovani/aula/infra/config"
)

func main() {
	err := config.InitConfig()
	CheckFatal(err)

	err = api.NewService().Start()
	CheckFatal(err)
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
