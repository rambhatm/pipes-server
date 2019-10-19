package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//Heroku mongo DB name
	herokuDB = "heroku_q5g3hk72"
	//Collections
	pipeCollection = "pipes"
)

var mongodbURI = os.Getenv("MONGODB_URI") + "?retryWrites=false"
var clientOptions = options.Client().ApplyURI(mongodbURI)

type pipe struct {
	user string `json:"user" bson:"user"`
	node string `json:"node" bson:"node"`
	data string `json:"data" bson:"data"`
}

func SetPipe(user string, node string, data string) {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("cannot connect to mongodb")
	}
	defer client.Disconnect(context.TODO())

	pipes := client.Database(herokuDB).Collection(pipeCollection)

	_, err = pipes.InsertOne(context.TODO(), pipe{user: user, node: node, data: data})
	if err != nil {
		log.Printf("error: unable to insert to user DB %s", err)
		return
	}

	return
}

func GetPipe(user string, node string) (myPipes []pipe) {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("cannot connect to mongodb")
	}
	defer client.Disconnect(context.TODO())

	pipes := client.Database(herokuDB).Collection(pipeCollection)
	matchingUserAndNode := bson.D{{"user", user}, {"node", node}}

	cur, err := pipes.Find(context.TODO(), matchingUserAndNode)
	if err != nil {
		log.Printf("Error on Finding all the documents %s", err)
	}
	for cur.Next(context.TODO()) {
		var p pipe
		err = cur.Decode(&p)
		if err != nil {
			log.Printf("Error on Decoding the document %s", err)
		}
		myPipes = append(myPipes, p)
	}
	return
}
