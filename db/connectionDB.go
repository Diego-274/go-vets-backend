package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx context.Context = context.TODO()
var DB *mongo.Database
var clientOptions = options.Client().ApplyURI("mongodb+srv://uo284016:password_dam@clustergo.bjil1.mongodb.net/?retryWrites=true&w=majority&appName=ClusterGo")

func ConnectDB() {
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		DB = client.Database("go-vets-backend")
		log.Println("Successful connection")
	}
}
