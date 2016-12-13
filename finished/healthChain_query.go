package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// returns the patient's data
// takes in 1 argument, patientID
func (t *HealthContract) getPatientInfo(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("getPatientInfo: You need to pass 1 argument")
	}

	parameters, err := myTableHandler.getPatient(stub, args[0])

	if err != nil {
		return nil, err
	}

	var p Patient
	p.ID = parameters[0]
	p.FirstName = parameters[1]
	p.LastName = parameters[2]
	p.Address = parameters[3]
	p.DateOfBirth = parameters[4]

	return json.Marshal(p)
}

// returns the permissions of a patient
// takes in 1 argument, patientID
func (t *HealthContract) getPermissionsOfPatient(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("getPermissionsOfPatient: You need to pass 1 argument")
	}

	return myTableHandler.getPatientPermissions(stub, args[0])
}

// returns all the patient's data (array of patient structs) of a doctor
// takes in 1 argument, doctorID
func (t *HealthContract) getPermissionsOfDoctor(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("getPermissionsOfDoctor: You need to pass 1 argument")
	}

	return myTableHandler.getDoctorPermissions(stub, args[0])
}

// returns the doctor's data
// takes in 1 argument, doctorID
func (t *HealthContract) getDoctorInfo(stub shim.ChaincodeStubInterface, args []string, caller string) ([]byte, error) {
	if caller == args[0] {
		fmt.Println("caller = patientID")
	}

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("getDoctorInfo: You need to pass 1 argument")
	}

	parameters, err := myTableHandler.getDoctor(stub, args[0])
	if err != nil {
		return nil, err
	}
	var d Doctor
	d.ID = parameters[0]
	d.FirstName = parameters[1]
	d.LastName = parameters[2]
	d.Address = parameters[3]
	d.DateOfBirth = parameters[4]
	d.Type = parameters[5]
	d.Institute = parameters[6]

	return json.Marshal(d)
}
