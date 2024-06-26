package backendrs

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

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

func InsertPasien(firstName string, lastName string, gender string, ttl string, status string, phonenumber string, alamat Alamat, medicalRecord MedicalRecord) (InsertedID interface{}) {
	var pasien Pasien
	pasien.FirstName = firstName
	pasien.LastName = lastName
	pasien.Gender = gender
	pasien.TTL = ttl
	pasien.Status = status
	pasien.Phone_number = phonenumber
	pasien.Alamat = alamat
	pasien.MedicalRecord = medicalRecord
	return InsertOneDoc("rumahsakit", "pasien", pasien)
}

func GetPasienFromPhoneNumber(phonenumber string) (data Pasien) {
	pasien := MongoConnect("rumahsakit").Collection("pasien")
	filter := bson.M{"phonenumber": phonenumber}
	err := pasien.FindOne(context.TODO(), filter).Decode(&pasien)
	if err != nil {
		fmt.Printf("getPasienFromPhoneNumber: %v\n", err)
	}
	return data
}

func GetAllPasien() (data []Pasien) {
	pasien := MongoConnect("rumahsakit").Collection("pasien")
	filter := bson.M{}
	cursor, err := pasien.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
