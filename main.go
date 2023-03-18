package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the placeholders with your credentials
const uri = "mongodb+srv://piyush:cknvQfIqOTSDMjdt@mongogo.otvbfso.mongodb.net/?retryWrites=true&w=majority"

func main() {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		cancel()
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	//// Send a ping to confirm a successful connection
	//var result bson.M
	//if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	//	panic(err)
	//}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	databases, err := client.ListDatabases(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("List of databases in your cluster: %v\n", databases)
}
