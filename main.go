package main

import (
	"context"
	"log"

	"github.com/piovani/aula/api"
	"github.com/piovani/aula/infra/config"
	"github.com/piovani/aula/infra/database"
	"github.com/piovani/aula/infra/database/mongo"
	"github.com/piovani/aula/infra/database/mongo/repositories"
)

func main() {
	var err error
	ctx := context.TODO()

	err = config.StarConfig()
	FatalError(err)

	db := GetDatabase(ctx)

	err = api.NewService(db).Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	studentRepository := repositories.NewStudentRepository(ctx, client)

	return database.NewDatabase(client, studentRepository)
}
