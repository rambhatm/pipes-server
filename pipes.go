package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//Heroku mongo DB name
	herokuDB = "heroku_q5g3hk72"
	//Collections
	pipes = "pipes"
)

var mongodbURI = os.Getenv("MONGODB_URI") + "?retryWrites=false"
var clientOptions = options.Client().ApplyURI(mongodbURI)

type pipe struct {
	user string
	node string
	data string
}

func SetPipe(user string, node string, data string) {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("cannot connect to mongodb")
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database(herokuDB).Collection(pipes)

	_, err = collection.InsertOne(context.TODO(), pipe{user: user, node: node, data: data})
	if err != nil {
		log.Printf("error: unable to insert to user DB %s", err)
		return
	}

	return
}

func GetPipe(user string) (myPipes []pipe) {
	return
}
