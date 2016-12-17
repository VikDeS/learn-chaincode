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

	patient, err := t.createPatient(stub, args, caller)

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

	doctor, err := t.createDoctor(stub, args, caller)

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

	arguments := []string{caller, args[0]}

	permission, err := t.createPermission(stub, arguments)

	// check if error during parsing of arguments
	if err != nil {
		return nil, err
	}

	// addPermission to table
	return myTableHandler.insertPermission(stub, permission)

}

// changes the values of row of Patient table with patientID = patient.patientID
// takes in 1 argument, Patient struct
func (t *HealthContract) changePatient(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	patient, err := t.createPatient(stub, args, caller)

	// check if error
	if err != nil {
		return nil, err
	}

	return myTableHandler.updatePatient(stub, patient)
}

// changes the values of row of Doctor table with doctorID = doctor.doctorID
// takes in 1 argument, Doctor struct
func (t *HealthContract) changeDoctor(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	doctor, err := t.createDoctor(stub, args, caller)

	// check if error
	if err != nil {
		return nil, err
	}

	return myTableHandler.updateDoctor(stub, doctor)
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

	return myTableHandler.deletePatient(stub, args[0])
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

	return myTableHandler.deleteDoctor(stub, args[0])
}

// removes permission to doctor to view patient's data
// takes in 2 arguments, PatientID and DoctorID
func (t *HealthContract) removePermission(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {

	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	arguments := []string{caller, args[0]}

	permission, err := t.createPermission(stub, arguments)

	// check if error
	if err != nil {
		return nil, err
	}

	return myTableHandler.deletePermission(stub, permission)
}

/*
----------------------------------------------------
------------------ HELPER METHODS ------------------
----------------------------------------------------
*/

// create patient object from arguments
func (t *HealthContract) createPatient(stub shim.ChaincodeStubInterface, args []string, caller string) (Patient, error) {
	var patient Patient

	// check number of arguments
	if len(args) != 5 {
		return patient, errors.New("createPatient: You need to pass 5 arguments")
	}
	// check if id is not empty
	if args[0] == "" {
		return patient, errors.New("createPatient: patientID may not be empty")
	}

	if caller != args[0] {
		fmt.Println("createPatient: caller != patientID")
	}

	// create Patient JSON
	patientID := "\"iD\":\"" + args[0] + "\", "
	firstname := "\"firstName\":\"" + args[1] + "\", "
	lastname := "\"lastName\":\"" + args[2] + "\", "
	address := "\"address\":\"" + args[3] + "\", "
	dateofbirth := "\"dateOfBirth\":\"" + args[4] + "\""

	patientJSON := "{" + patientID + firstname + lastname + address + dateofbirth + "}"

	// parse Patient JSON to patient go struct
	err := json.Unmarshal([]byte(patientJSON), &patient)

	return patient, err
}

// create doctor object from arguments
func (t *HealthContract) createDoctor(stub shim.ChaincodeStubInterface, args []string, caller string) (Doctor, error) {
	var doctor Doctor

	// check number of arguments
	if len(args) != 7 {
		return doctor, errors.New("createDoctor: You need to pass 7 arguments")
	}
	// check if id is not empty
	if args[0] == "" {
		return doctor, errors.New("createDoctor: doctorID may not be empty")
	}

	if caller == args[0] {
		fmt.Println("createDoctor: caller = doctorID")
	}

	// create Doctor JSON
	doctorID := "\"iD\":\"" + args[0] + "\", "
	firstname := "\"firstName\":\"" + args[1] + "\", "
	lastname := "\"lastName\":\"" + args[2] + "\", "
	address := "\"address\":\"" + args[3] + "\", "
	dateofbirth := "\"dateOfBirth\":\"" + args[4] + "\", "
	doctorType := "\"type\":\"" + args[5] + "\", "
	institute := "\"institute\":\"" + args[6] + "\""

	doctorJSON := "{" + doctorID + firstname + lastname + address + dateofbirth + doctorType + institute + "}"

	// parse Doctor JSON to doctor go struct
	err := json.Unmarshal([]byte(doctorJSON), &doctor)

	return doctor, err
}

// create doctor object from arguments
func (t *HealthContract) createPermission(stub shim.ChaincodeStubInterface, args []string) (Permission, error) {
	var permission Permission

	// check number of arguments
	if len(args) != 2 {
		return permission, errors.New("createPermission: You need to pass 2 argument")
	}

	// create Patient JSON
	patientID := "\"iD\":\"" + args[0] + "\", "
	doctorID := "\"firstName\":\"" + args[1] + "\", "

	permissionJSON := "{" + patientID + doctorID + "}"

	// parse Patient JSON to patient go struct
	err := json.Unmarshal([]byte(permissionJSON), &permission)

	return permission, err
}
