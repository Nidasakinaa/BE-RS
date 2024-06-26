package backendrs

import (
	"fmt"
	"testing"
)

func TestInsertPasien(t *testing.T) {
	firstName := "Nisa"
	lastName := "Karina"
	gender := "Perempuan"
	ttl := "Bandung, 12 Januari 1990"
	status := "Belum Menikah"
	phonenumber := "089765"
	alamat := Alamat{
		Jalan:   "Aceh",
		Kota:    "Banda Aceh",
		KodePos: "71403",
	}
	medicalRecord := MedicalRecord{
		VisitDate:  "12 Maret 2023",
		DoctorName: "Andrew",
		Diagnosis:  "Stroke",
		Treatment:  "Carotid endarterectomy",
		Notes:      "-",
	}
	hasil := InsertPasien(firstName, lastName, gender, ttl, status, phonenumber, alamat, medicalRecord)
	fmt.Println(hasil)
}

func TestGetPasienFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	pasien := GetPasienFromPhoneNumber(phonenumber)
	fmt.Println(pasien)
}

func TestGetAll(t *testing.T) {
	data := GetAllPasien()
	fmt.Println(data)
}
