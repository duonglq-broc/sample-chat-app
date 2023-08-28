package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
}

func InitDB(ctx context.Context, host string) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoDB{
		Client: client,
	}, err
}

func (m *MongoDB) Close(ctx context.Context) error {
	err := m.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
