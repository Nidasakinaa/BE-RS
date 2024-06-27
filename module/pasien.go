package module

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPasien(db *mongo.Database, col string, firstName string, lastName string, gender string, ttl string, status string, phonenumber string) (insertedID primitive.ObjectID, err error) {
	pasien := bson.M{
		"firstName":   firstName,
		"lastName":    lastName,
		"gender":      gender,
		"ttl":         ttl,
		"status":      status,
		"phonenumber": phonenumber,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), pasien)
	if err != nil {
		fmt.Printf("InsertPasien: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
