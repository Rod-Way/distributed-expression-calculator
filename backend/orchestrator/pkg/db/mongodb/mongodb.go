package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
	Ctx    context.Context
}

func New(uri, dbName string, timeout time.Duration) (*Mongo, error) {
	// Creating context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// init client
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	// Check connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Mongo{
		Client: client,
		DB:     client.Database(dbName),
		Ctx:    context.Background(),
	}, nil
}

func (m *Mongo) Disconnect() error {
	return m.Client.Disconnect(m.Ctx)
}
