package cmd

import (
	"context"

	"github.com/piovani/aula/api"
	"github.com/piovani/aula/infra/cache"
	"github.com/piovani/aula/infra/config"
	"github.com/spf13/cobra"
)

var Api = &cobra.Command{
	Use:   "api",
	Short: "API manager student register",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ctx := context.TODO()

		err = config.StarConfig()
		FatalError(err)

		db := GetDatabase(ctx)
		cache := cache.NewCache()

		err = api.NewService(db, cache).Start()
		FatalError(err)
	},
}
