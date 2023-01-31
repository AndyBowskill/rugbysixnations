package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type database func() teamDatabase

type teamDatabase struct {
	ID        primitive.ObjectID `bson:"_id"`
	HeadCoach string
	Name      string
	Forwards  interface{}
	Backs     interface{}
}
