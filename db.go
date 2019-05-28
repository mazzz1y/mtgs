package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo" 
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoHost = "127.0.0.1"
	mongoPort = "27017"
	mongoDB = "mtproto"
)

func InitDB() *mongo.Database {
	uri := "mongodb://" + mongoHost + ":" + mongoPort
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	if client.Ping(ctx, readpref.Primary()) != nil {
		panic(err)
	}
	return client.Database(mongoDB)
}