package backendrs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alamat struct {
	Jalan   string `bson:"jalan,omitempty" json:"jalan,omitempty"`
	Kota    string `bson:"kota,omitempty" json:"kota,omitempty"`
	KodePos string `bson:"kodepos,omitempty" json:"kodepos,omitempty"`
}

type MedicalRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	VisitDate  string             `bson:"visitdate,omitempty" json:"visitdate,omitempty"`
	DoctorName string             `bson:"doctor,omitempty" json:"doctor,omitempty"`
	Diagnosis  string             `bson:"diagnosis,omitempty" json:"diagnosa,omitempty"`
	Treatment  string             `bson:"treatment,omitempty" json:"treatment,omitempty"`
	Notes      string             `bson:"notes,omitempty" json:"notes,omitempty"`
}

type Pasien struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName     string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName      string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	TTL           string             `bson:"ttl,omitempty" json:"ttl,omitempty"`
	Status        string             `bson:"status,omitempty" json:"status,omitempty"`
	Phone_number  string             `bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
	Alamat        Alamat             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	MedicalRecord MedicalRecord      `bson:"medicalRecord,omitempty" json:"medicalRecord,omitempty"`
}
