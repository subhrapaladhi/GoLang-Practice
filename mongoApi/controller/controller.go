package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")

	// collection instance
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("mongo collection is ready")
}
