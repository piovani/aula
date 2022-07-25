package mongo

import (
	"context"

	"github.com/piovani/aula/infra/config"
)

func NweClient(ctx context.context) (client, err) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*config.Env.MongoTimeout)
	defer cancel()

	options := options.Client().ApplyURI(config.Env.MongoUrl)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return client, err
	}

	if err = client.Ping(); err != nil {
		return client, err
	}

	return client, nil
}
