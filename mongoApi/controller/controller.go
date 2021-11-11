package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/subhrapaladhi/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Mongodb helpers

func insertOneMovie(movie model.Netflix) {
	data, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	data, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated count = ", data.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count = ", deleteCount)
}

func deleteAllMovies() {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count = ", deleteResult.DeletedCount)
}

func getAllMoview() {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	fmt.Println(movies)
}
