package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dburi string
)

func init() {
	dburi = os.Getenv("MONGO_URL")
}

func GetClient() (*mongo.Client, context.Context, func()) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(dburi).SetServerAPIOptions(serverAPIOptions)

	// get context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, func() {
		client.Disconnect(ctx)
		cancel()
	}
}

func Connect() {
	client, ctx, cancel := GetClient()
	defer cancel()
	// ping
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("while pinging: ", err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
