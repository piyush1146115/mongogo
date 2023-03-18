package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

	quickStartDB := client.Database("quickstart")
	podcastCollection := quickStartDB.Collection("podcasts")
	//podcastResult, err := podcastCollection.InsertOne(ctx, bson.D{
	//	{Key: "title", Value: "The polyglot developer"},
	//	{Key: "author", Value: "Piyush"},
	//	{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(podcastResult.InsertedID)
	//
	episodesCollection := quickStartDB.Collection("episodes")
	//episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
	//	bson.D{
	//		{"podcast", podcastResult.InsertedID},
	//		{"title", "Episode #1"},
	//		{"description", "This is the first episode"},
	//		{"duration", 25},
	//	},
	//	bson.D{
	//		{"podcast", podcastResult.InsertedID},
	//		{"title", "Episode #2"},
	//		{"description", "This is the second episode"},
	//		{"duration", 32},
	//	},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(episodeResult.InsertedIDs)

	//cursor, err := episodesCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var episodes []bson.M
	//if err = cursor.All(ctx, &episodes); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(episodes)

	// Get all the elements of episodes collection
	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		//fmt.Println(episode)
	}

	// Get one element of PodCase collection
	var podcast bson.M
	if err = podcastCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(podcast)

	// Get episodes whose duration == 25
	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(episodesFiltered)

	// sort episodes and filter it based on the duration field
	opt := options.Find()
	opt.SetSort(bson.D{{"duration", 1}})
	sortCursor, err := episodesCollection.Find(ctx, bson.D{
		{"duration", bson.D{
			{"$gt", 24},
		}},
	}, opt)

	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)
}
