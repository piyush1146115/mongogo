package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func createData() {
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

	quickStartDB := client.Database("quickstart")
	podcastCollection := quickStartDB.Collection("podcasts")

	podcastResult, err := podcastCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The polyglot developer"},
		{Key: "author", Value: "Piyush"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(podcastResult.InsertedID)

	episodesCollection := quickStartDB.Collection("episodes")
	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode #1"},
			{"description", "This is the first episode"},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode #2"},
			{"description", "This is the second episode"},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodeResult.InsertedIDs)
}
