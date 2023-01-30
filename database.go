package main

import "go.mongodb.org/mongo-driver/bson"

type database func() bson.M
