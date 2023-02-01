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

func getTeamFromMongoDB(teamName string) teamDatabase {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	uri := os.Getenv("MONGODB_URI")

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("rugbysixnations").Collection("teams")

	var team teamDatabase
	err = collection.FindOne(ctx, bson.D{{"name", teamName}}).Decode(&team)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the team name %s\n", teamName)
	}
	if err != nil {
		panic(err)
	}

	return team
}
