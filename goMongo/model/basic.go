package model

import (
	"context"
	"log"
	"mongodb_demo/connection"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "Info"

type Parent struct {
	Name string
	Age  int
}
type Info struct {
	Name    *string
	Age     *int
	Male    *bool
	Hobby   []string
	Parents []Parent
}

func InfoCreateCollection() {
	db, ctx := connection.GetDB()
	err := db.CreateCollection(ctx, collectionName, options.CreateCollection())
	if err != nil {
		log.Fatalln(err)
	}
}
func InfoGetCollection() (*mongo.Collection, context.Context) {
	db, ctx := connection.GetDB()
	collection := db.Collection(collectionName, options.Collection())
	return collection, ctx
}

func InfoGetCollectionDrop() {
	colletion, ctx := InfoGetCollection()
	err := colletion.Drop(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
