package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// adds the patient to the Patients table
// needs 5 arguments: patient id, firstname, lastname, address, dateofbirth
func (t *HealthContract) addPatient(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	var patient Patient

	// check number of arguments
	if len(args) != 5 {
		return nil, errors.New("addPatient: You need to pass 5 arguments")
	}
	// check if id is not empty
	if args[0] == "" {
		return nil, errors.New("addPatient: patientID may not be empty")
	}

	if caller == args[0] {
		fmt.Println("addPatient: caller = patientID")
	}

	// create Patient JSON
	patient_id := "\"iD\":\"" + args[0] + "\", "
	firstname := "\"firstName\":\"" + args[1] + "\", "
	lastname := "\"lastName\":\"" + args[2] + "\", "
	address := "\"address\":\"" + args[3] + "\", "
	dateofbirth := "\"dateOfBirth\":\"" + args[4] + "\""

	patient_json := "{" + patient_id + firstname + lastname + address + dateofbirth + "}"

	// parse Patient JSON to patient go struct
	err := json.Unmarshal([]byte(patient_json), &patient)

	// check if error during parsing of arguments
	if err != nil {
		return nil, err
	}

	// addPatient to table
	return myTableHandler.insertPatient(stub, patient)

}

// calls authenticator to register, enroll new doctor and
// adds the doctor struct to the Doctor table
// needs 2 arguments: Doctor struct and password
func (t *HealthContract) addDoctor(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	var doctor Doctor

	// check number of arguments
	if len(args) != 7 {
		return nil, errors.New("addDoctor: You need to pass 7 arguments")
	}
	// check if id is not empty
	if args[0] == "" {
		return nil, errors.New("addDoctor: doctorID may not be empty")
	}

	if caller == args[0] {
		fmt.Println("addDoctor: caller = doctorID")
	}

	// create Doctor JSON
	doctor_id := "\"iD\":\"" + args[0] + "\", "
	firstname := "\"firstName\":\"" + args[1] + "\", "
	lastname := "\"lastName\":\"" + args[2] + "\", "
	address := "\"address\":\"" + args[3] + "\", "
	dateofbirth := "\"dateOfBirth\":\"" + args[4] + "\", "
	doctor_type := "\"type\":\"" + args[5] + "\", "
	institute := "\"institute\":\"" + args[6] + "\""

	doctor_json := "{" + doctor_id + firstname + lastname + address + dateofbirth + doctor_type + institute + "}"

	// parse Doctor JSON to doctor go struct
	err := json.Unmarshal([]byte(doctor_json), &doctor)

	// check if error during parsing of arguments
	if err != nil {
		return nil, err
	}

	// addDoctor to table
	return myTableHandler.insertDoctor(stub, doctor)

}

// gives permission to doctor to view patient's data
// takes in 1 arguments, PatientID and DoctorID
func (t *HealthContract) giveDocorPermission(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	var permission Permission

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("giveDocorPermission: You need to pass 1 argument")
	}

	// create Patient JSON
	patient_id := "\"iD\":\"" + caller + "\", "
	doctor_id := "\"firstName\":\"" + args[0] + "\", "

	permission_json := "{" + patient_id + doctor_id + "}"

	// parse Patient JSON to patient go struct
	err := json.Unmarshal([]byte(permission_json), &permission)

	// check if error during parsing of arguments
	if err != nil {
		return nil, err
	}

	// addPermission to table
	return myTableHandler.insertPermission(stub, permission)

}

// NOG TE IMPLEMENTEREN FUNCTIES

// changes the values of row of Patient table with patientID = patient.patientID
// takes in 1 argument, Patient struct
func (t *HealthContract) changePatient(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("changePatient: You need to pass 1 argument")
	}

	return nil, nil
}

// changes the values of row of Doctor table with doctorID = doctor.doctorID
// takes in 1 argument, Doctor struct
func (t *HealthContract) changeDoctor(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("changeDoctor: You need to pass 1 argument")
	}

	return nil, nil
}

// removes the patient out of Patient table with patientID = patient.patientID
// takes in 1 argument, Patient struct
func (t *HealthContract) removePatient(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("removePatient: You need to pass 1 argument")
	}

	return nil, nil
}

// removes the doctor out of Doctor table with doctorID = doctor.doctorID
// takes in 1 argument, Doctor struct
func (t *HealthContract) removeDoctor(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("removeDoctor: You need to pass 1 argument")
	}

	return nil, nil
}

// removes permission to doctor to view patient's data
// takes in 2 arguments, PatientID and DoctorID
func (t *HealthContract) removePermission(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("removePermission: You need to pass 2 arguments")
	}

	return nil, nil
}
