package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewDb(dbConfig *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	connectionUri := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri))
	if err != nil {
		return nil, err
	}
	return client.Database(dbConfig.DbName), nil
}
