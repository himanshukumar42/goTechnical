package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient() (*mongo.Client, error) {
	ctx, cancel :=  context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	host := "localhost"
	port := "27017"
	db_url := fmt.Sprintf("mongodb://%v:%v", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))
	if err != nil {
		return nil, err
	}
	return client, nil
}