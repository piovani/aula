package main

import (
	"github.com/piovani/aula/api"
)

func main() {
	service := api.NewService()
	service.Start()
}
