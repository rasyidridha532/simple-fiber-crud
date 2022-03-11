package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"simple-fiber-crud/exception"
	"time"
)

func Connect() (*mongo.Database, error) {
	ctx, _ := NewMongoContext()

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connURI := "mongodb://%v:%v@%v:%v"
	if dbUser == "" {
		connURI = "mongodb://%s%s%v:%v"
	}

	clientURI := fmt.Sprintf(connURI,
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
	)

	clientOptions := options.Client().
		ApplyURI(clientURI).
		SetMinPoolSize(10).
		SetMaxPoolSize(30).SetMaxConnIdleTime(time.Duration(5) * time.Second)
	client, errors := mongo.NewClient(clientOptions)
	exception.LogError(errors)

	errors = client.Connect(ctx)
	exception.LogError(errors)

	return client.Database(dbName), errors
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
