package storage

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client *mongo.Client
	DB     *mongo.Database
}

func (m *MongoStorage) Close() error {
	return m.client.Disconnect(context.Background())
}

func newMongoStorage(client *mongo.Client, db *mongo.Database) *MongoStorage {
	return &MongoStorage{client: client, DB: db}
}

var ErrDBNameRequired = errors.New("database name is required")

func BootstrapMongo(uri, dbName string, timeout time.Duration) (*MongoStorage, error) {
	if dbName == "" {
		return nil, ErrDBNameRequired
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	return newMongoStorage(client, client.Database(dbName)), nil
}
