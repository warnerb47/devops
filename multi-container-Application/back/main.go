package main

import (
	"context"
	"fmt"
	"log"
	"warnerb47/todo/controllers"
	"warnerb47/todo/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	server         *gin.Engine
	todoService    services.TodoService
	todoController *controllers.TodoController
	ctx            context.Context
	todoCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	API_URI := "mongodb://root:root@localhost:27017"
	mongoConnection := options.Client().ApplyURI(API_URI)
	mongoClient, err = mongo.Connect(mongoConnection)
	if err != nil {
		log.Fatal(err)
	}
	err := mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	todoCollection = mongoClient.Database("todoDB").Collection("todos")
	todoService = services.New(todoCollection, ctx)
	todoController = controllers.New(todoService)
	server = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	server.Use(cors.New(config))
}

func main() {
	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("/v1")
	todoController.RegisterTodoRoutes(basePath)
	log.Fatal(server.Run(":8080"))
}
