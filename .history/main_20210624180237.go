package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//connecting the mongo server to golang
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

	//disconnecting the mongo server after operation
    defer client.Disconnect(ctx)
    fmt.Println("connected")
	
	//creating the database
    quickstartDatabase := client.Database("testuser")

	//creating the collection
    podcastsCollection := quickstartDatabase.Collection("podcasts")
 
	//inserating the single document
	//bson stands for binary form of json stored in mongodb
	//bson.D{stands for bson document}, bson.M{stands for bson map}, bson.A{stands for bson array}
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
    {Key: "title", Value: "Developer"},
    {Key: "author",  Value: "abc"},
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(podcastResult.InsertedID)

//finding the document or reading the document
var podcast bson.M
err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast)
if err != nil {
    log.Fatal(err)
}
fmt.Println(podcast)

//updating the document fileds
result, err := podcastsCollection.ReplaceOne(
    ctx,
    bson.M{"author": "abc"},
    bson.M{
        "title":  "Tester",
        "author": "xyz",
    },
)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Replaced %v Documents!\n", result.ModifiedCount)

//deleting the document
// result, err := podcastsCollection.DeleteOne(ctx, bson.M{"title": "Developer"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("DeleteOne removed %v document(s)\n", result)
}