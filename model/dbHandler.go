package model

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

type dbHandler struct {
	db   *mongo.Client
	coll *mongo.Collection
}

var mgo dbHandler

func addDB(data *TodoData) bool {
	insertResult, err := mgo.coll.InsertOne(context.TODO(), *data)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println("MongoDB: Added Data: ", insertResult)
	return true
}
func getSize() int {
	count, err := mgo.coll.EstimatedDocumentCount(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return int(count)
}
func getDB() []*TodoData {
	list := []*TodoData{}

	filter := bson.D{{}}
	cursor, err := mgo.coll.Find(context.TODO(), filter.Map())
	if err != nil {
		log.Fatal("MongoDB : getDB ", err)
	}

	for cursor.Next(context.TODO()) {
		var temp TodoData
		err := cursor.Decode(&temp)
		if err != nil {
			log.Fatal("MongoDB : getDB Decode ", err)
		}
		list = append(list, &temp)
	}
	return list
}
func deleteDB(id int) bool {
	dResult, err := mgo.coll.DeleteOne(context.TODO(), bson.D{{"id", id}}.Map())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MongoDB : Delete Result : ", dResult)
	return true
}

func NewDB() {
	mgo.db, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:10102"))
	err := mgo.db.Connect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Conneted MongoDB....")
	mgo.coll = mgo.db.Database("todosReact").Collection("todo")
}

func CloseDB() {
	err := mgo.db.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected by MongoDB....")
}
