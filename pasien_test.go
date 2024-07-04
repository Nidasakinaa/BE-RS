package backendrs

import (
	"fmt"
	"testing"

	model "github.com/Nidasakinaa/BackendRS/model"
	module "github.com/Nidasakinaa/BackendRS/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetPasienByID(t *testing.T) {
	_id := "667e6038323e1ea07fdf3340"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPasienByID(objectID, module.MongoConn, "pasien")
	if err != nil {
		t.Fatalf("error calling GetPasienFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllPasien(module.MongoConn, "pasien")
	fmt.Println(data)
}

func TestInsertPasien(t *testing.T) {
	pasienName := "Naya Kanaya"
	gender := "Perempuan"
	ttl := "Bandung, 20 Mei 2006"
	status := "Belum Menikah"
	phonenumber := "13456"
	alamat := "Jl.Bandung, Gorontalo, 71903"
	doctor := model.Doctor{
		Name:      "Andre",
		Specialty: "Oncology",
		Contact:   "123-456-7890",
	}
	medicalRecord := model.MedicalRecord{
		VisitDate:  "12 Juli 2023",
		DoctorName: "Andre",
		Diagnosis:  "Cancer",
		Treatment:  "Kemo",
		Notes:      "-",
	}
	insertedID, err := module.InsertPasien(module.MongoConn, "pasien", pasienName, gender, ttl, status, phonenumber, alamat, doctor, medicalRecord)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

// func TestUpdatePasien(t *testing.T) {
// 	id := "66756b2bb0584a360d9709d8"
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("Error converting id to ObjectID: %v", err)
// 	}

// 	pasienName := "Nanad Nadia"
// 	gender := "Perempuan"
// 	ttl := "Bandung, 20 Juli 1998"
// 	status := "Belum Menikah"
// 	phonenumber := "12345"
// 	alamat := "Bandung, Bandung, 71903"
// 	doctor := model.Doctor{
// 		Name:      "Andre",
// 		Specialty: "Oncology",
// 		Contact:   "123-456-7890",
// 	}
// 	medicalRecord := model.MedicalRecord{
// 		VisitDate:  "12 Juli 2023",
// 		DoctorName: "Andre",
// 		Diagnosis:  "Cancer",
// 		Treatment:  "Kemo",
// 		Notes:      "-",
// 	}
// 	err = module.UpdatePasien(context.TODO(), module.MongoConn, "pasien", objectID, pasienName, gender, ttl, status, phonenumber, alamat, doctor, medicalRecord)
// 	if err != nil {
// 		t.Fatalf("Error calling UpdatePasien: %v", err)
// 	}
// 	fmt.Println("Data berhasil diupdate")
// }

// func TestDeletePasienByID(t *testing.T) {
// 	id := "667c07111601cc599e30ade2" // ID data yang ingin dihapus
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeletePasienByID(objectID, module.MongoConn, "pasien")
// 	if err != nil {
// 		t.Fatalf("error calling TestDeletePasienByID: %v", err)
// 	}

// 	_, err = module.GetPasienByID(objectID, module.MongoConn, "pasien")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }
