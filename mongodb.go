package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getTeamFromMongoDB(teamName string) bson.M {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("rugbysixnations").Collection("teams")

	var team bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"name", teamName}}).Decode(&team)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the team name %s\n", teamName)
		return nil
	}
	if err != nil {
		panic(err)
	}

	return team
}
