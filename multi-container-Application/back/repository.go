package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var todoCollection *mongo.Collection
var API_URI = "mongodb://foo:bar@localhost:27017"

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(API_URI))
	if err != nil {
		log.Fatal("Error while connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Can not ping MongoDB:", err)
	}

	todoCollection = client.Database("todo_db").Collection("todos")
	log.Println("Connecté à MongoDB !")
}
