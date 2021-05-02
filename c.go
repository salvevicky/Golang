package main

import ("fmt"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))//your ip address and mongodb port number
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    datab := client.Database("Vikson")// database name quicstart
    coll1 := datab.Collection("info")// collection name info
    //coll2 := datab.Collection("episodes")
	// inserting data into collection name info
	podcastResult, err := coll1.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	fmt.Println(podcastResult)
}