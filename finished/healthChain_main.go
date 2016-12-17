package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var myTableHandler = NewTableHandler()

// User has common info
type User struct {
	ID          string `json:"iD"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
}

// Patient has common info
type Patient struct {
	User
}

// Doctor has more specific info attached
type Doctor struct {
	User
	Type      string `json:"type"`
	Institute string `json:"institute"`
}

// Permission to see eachothers data
type Permission struct {
	PatientID string `json:"patientID"`
	DoctorID  string `json:"doctorID"`
}

// HealthContract structure
type HealthContract struct {
}

/*
---------------------------------------------------
------------------ DEPLOY/INIT --------------------
---------------------------------------------------
*/
// main function
func main() {
	err := shim.Start(new(HealthContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

//Init the chaincode asigned the value "0" to the counter in the state.
func (t *HealthContract) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return myTableHandler.createTables(stub)
}

/*
---------------------------------------------------
--------------- INVOKE FUNCTIONALITY --------------
---------------------------------------------------
*/

//Invoke Transaction makes increment counter
func (t *HealthContract) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	caller, typeOfUser, err := t.getCallerData(stub)

	if err != nil {
		return nil, err
	}

	switch typeOfUser {

	/*
		-----------------------------------------
		--------- PATIENT FUNCTIONALITY ---------
		-----------------------------------------
	*/
	case "PATIENT":
		switch function {

		case "changeMyInfo":
			return t.changePatient(stub, args, caller)
		case "removeMyAccount":
			return t.removePatient(stub, args, caller)
		case "givePermission":
			return t.giveDocorPermission(stub, args, caller)
		case "removePermission":
			return t.removePermission(stub, args, caller)
		default:
			return nil, errors.New("Invalid invoke function name for Patient")
		}

		/*
			-----------------------------------------
			---------- DOCTOR FUNCTIONALITY ---------
			-----------------------------------------
		*/
	case "DOCTOR":
		switch function {

		case "changePatientInfo":
			return t.changePatient(stub, args, caller)
		case "changeMyInfo":
			return t.changeDoctor(stub, args, caller)
		case "removeMyAccount":
			return t.removeDoctor(stub, args, caller)
		default:
			return nil, errors.New("Invalid invoke function name for Doctor")
		}

		/*
			-----------------------------------------
			---------- ADMIN FUNCTIONALITY ----------
			-----------------------------------------
		*/
	case "ADMIN":
		switch function {

		case "addPatient":
			return t.addPatient(stub, args, caller)
		case "addDoctor":
			return t.addDoctor(stub, args, caller)
		default:
			return nil, errors.New("Invalid invoke function name for Admin")
		}

	default:
		return nil, errors.New("INVOKE: Invalid typeOfUser")
	}
}

/*
---------------------------------------------------
--------------- QUERY FUNCTIONALITY ---------------
---------------------------------------------------
*/

// Query callback representing the query of a chaincode
func (t *HealthContract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	caller, typeOfUser, err := t.getCallerData(stub)

	if err != nil {
		return nil, err
	}

	switch typeOfUser {
	/*
		-----------------------------------------
		--------- PATIENT FUNCTIONALITY ---------
		-----------------------------------------
	*/
	case "PATIENT":
		switch function {

		case "getPatientInfo": //case "getMyInfo":
			return t.getPatientInfo(stub, args, caller)
		case "getPermissions":
			return t.getPermissionsOfPatient(stub, args, caller)

		default:
			return nil, errors.New("Invalid invoke function name for Patient")
		}

		/*
			-----------------------------------------
			---------- DOCTOR FUNCTIONALITY ---------
			-----------------------------------------
		*/
	case "DOCTOR":
		switch function {

		case "getMyPermissions":
			return t.getPermissionsOfDoctor(stub, args, caller)
		case "getPatientInfo":
			return t.getPatientInfo(stub, args, caller)
		case "getMyInfo":
			return t.getDoctorInfo(stub, args, caller)
		default:
			return nil, errors.New("Invalid invoke function name for Doctor")
		}

		/*
			-----------------------------------------
			---------- ADMIN FUNCTIONALITY ----------
			-----------------------------------------
		*/
	case "ADMIN":
		switch function {

		default:
			return nil, errors.New("Invalid query function name for Admin")
		}

	default:
		return nil, errors.New("QUERY: Invalid typeOfUser")
	}
}
