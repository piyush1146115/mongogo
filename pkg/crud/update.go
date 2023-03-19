package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func updateData() {
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

	database := client.Database("quickstart")
	podcastCollection := database.Collection("podcasts")

	id, _ := primitive.ObjectIDFromHex("6415899a020cba0b02a482e1")

	result, err := podcastCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"author", "Piyush Kanti Das"}}},
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	result, err = podcastCollection.UpdateMany(
		ctx,
		bson.M{"title": "The polyglot developer"},
		bson.D{
			{"$set", bson.D{{"author", "Piyush Kanti Das"}}},
		},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	result, err = podcastCollection.ReplaceOne(
		ctx,
		bson.M{"author": "Piyush Kanti Das"},
		bson.M{
			"author": "Piyush",
			"title":  "The Piyush show",
		},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}
