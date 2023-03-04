package cmd

import (
	"context"
	"log"

	"github.com/piovani/aula/infra/database"
	"github.com/piovani/aula/infra/database/mongo"
	"github.com/piovani/aula/infra/database/mongo/repositories"
	"github.com/spf13/cobra"
)

func Execute() {
	root := &cobra.Command{
		Short:   "Student Mananeger",
		Version: "1.0.0",
	}

	root.AddCommand(
		Api,
	)

	root.Execute()
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
