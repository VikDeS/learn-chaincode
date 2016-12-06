package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var myTableHandler = NewTableHandler()

//var myAuthenticator = NewAuthenticator()

type User struct {
	ID          uint64 `json:"iD"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
}

type Patient struct {
	User
}

type Doctor struct {
	User
	Function  string `json:"function"`
	Institute string `json:"institute"`
}

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

	//switch role {

	/*
		-----------------------------------------
		--------- PATIENT FUNCTIONALITY ---------
		-----------------------------------------
	*/
	//case:Patient:
	switch function {

	case "changeMyInfo":
		return t.changePatient(stub, args)
	case "removeMyAccount":
		return t.removePatient(stub, args)
	case "givePermission":
		return t.givePermission(stub, args)
	case "removePermission":
		return t.removePermission(stub, args)

	/*default:
			return nil, errors.New("Invalid invoke function name for Patient.")
	}*/

	/*
		-----------------------------------------
		---------- DOCTOR FUNCTIONALITY ---------
		-----------------------------------------
	*/
	//case:Doctor:
	//switch function {
	case "changePatientInfo":
		return t.changePatient(stub, args)
	case "changeMyInfo":
		return t.changeDoctor(stub, args)
	case "removeMyAccount":
		return t.removeDoctor(stub, args)

	/*default:
				return nil, errors.New("Invalid invoke function name for Doctor.")
	}*/

	/*
		-----------------------------------------
		---------- ADMIN FUNCTIONALITY ----------
		-----------------------------------------
	*/
	//case:Admin:
	//switch function {
	case "addPatient":
		return t.addPatient(stub, args)
	case "addDoctor":
		return t.addDoctor(stub, args)
	default:
		return nil, errors.New("Invalid invoke function name for Admin.")
		//}

		/*default:
		return nil, errors.New("Invalid invoke function name for Doctor.")*/
	}
}

/*-------------------------------------------------------------------------------------------------------------*/

// changes the values of row of Patient table with patientID = patient.patientID
// takes in 1 argument, Patient struct
func (t *HealthContract) changePatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("changePatient: You need to pass 1 argument")
	}

	return nil, nil
}

// changes the values of row of Doctor table with doctorID = doctor.doctorID
// takes in 1 argument, Doctor struct
func (t *HealthContract) changeDoctorInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("changeDoctor: You need to pass 1 argument")
	}

	return nil, nil
}

// removes the patient out of Patient table with patientID = patient.patientID
// takes in 1 argument, Patient struct
func (t *HealthContract) removePatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("removePatient: You need to pass 1 argument")
	}

	return nil, nil
}

// removes the doctor out of Doctor table with doctorID = doctor.doctorID
// takes in 1 argument, Doctor struct
func (t *HealthContract) removeDoctor(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("removeDoctor: You need to pass 1 argument")
	}

	return nil, nil
}

// gives permission to doctor to view patient's data
// takes in 2 arguments, PatientID and DoctorID
func (t *HealthContract) givePermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 2 {
		return nil, errors.New("givePermission: You need to pass 2 arguments")
	}

	return nil, nil
}

// removes permission to doctor to view patient's data
// takes in 2 arguments, PatientID and DoctorID
func (t *HealthContract) removePermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("removePermission: You need to pass 2 arguments")
	}

	return nil, nil
}

// calls authenticator to register, enroll new patient and
// adds the patient struct to the Patients table
// needs 2 arguments: Patient struct and password
func (t *HealthContract) addPatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var pw string
	var patient Patient

	// check number of arguments
	if len(args) != 2 {
		return nil, errors.New("addPatient: You need to pass 2 arguments")
	}

	// parse first argument to patient object
	patient = args[0]

	// put second argument in password string
	pw = args[1]

	// check if error during parsing of arguments
	if err != nil {
		return nil, err
	}

	//var stream
	// register and enroll Patient
	//stream, err := myAuthenticator.enrollPatient(stub, patient.PatientID, pw)

	// check if error during enrolling
	if err != nil {
		return nil, err
	}
	// addPatient to table
	return nil, myTableHandler.insertPatient(stub, patient.PatientID, patient.FirstName, patient.LastName, patient.Address, patient.DateOfBirth)

}

// calls authenticator to register, enroll new doctor and
// adds the doctor struct to the Doctor table
// needs 2 arguments: Doctor struct and password
func (t *HealthContract) addDoctor(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var pw string
	var doctor Doctor

	// check number of arguments
	if len(args) != 2 {
		return nil, errors.New("addDoctor: You need to pass 2 arguments")
	}

	// parse first argument to doctor object
	doctor = args[0]

	// put second argument in password string
	pw = args[1]

	// check if error during parsing of arguments
	if err != nil {
		return nil, err
	}

	//var stream
	// register and enroll Doctor
	//stream, err := myAuthenticator.enrollDoctor(stub, doctor.doctorID, pw)

	// check if error during enrolling
	if err != nil {
		return nil, err
	}
	// addDoctor to table
	return nil, myTableHandler.insertDoctor(stub, doctor.DoctorID, doctor.FirstName, doctor.LastName, doctor.Function, doctor.Institute, doctor.Address, doctor.DateOfBirth)

}

/*
---------------------------------------------------
--------------- QUERY FUNCTIONALITY ---------------
---------------------------------------------------
*/

// Query callback representing the query of a chaincode
func (t *HealthContract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	//switch role {

	/*
		-----------------------------------------
		--------- PATIENT FUNCTIONALITY ---------
		-----------------------------------------
	*/
	//case:Patient:
	switch function {

	case "getMyInfo":
		return t.getPatientInfo(stub, args)
	case "getPermission":
		return t.getPermission(stub, args)

	/*default:
			return nil, errors.New("Invalid invoke function name for Patient.")
	}*/

	/*
		-----------------------------------------
		---------- DOCTOR FUNCTIONALITY ---------
		-----------------------------------------
	*/
	//case:Doctor:
	//switch function {
	case "getMyPatients":
		return t.getPatientsOfDoctor(stub, args)
	case "getPatientInfo":
		return t.getPatientInfo(stub, args)
	case "getMyInfo":
		return t.getDoctorInfo(stub, args)

	/*default:
				return nil, errors.New("Invalid invoke function name for Doctor.")
	}*/

	/*
		-----------------------------------------
		---------- ADMIN FUNCTIONALITY ----------
		-----------------------------------------
	*/
	//case:Admin:
	//switch function {

	default:
		return nil, errors.New("Invalid invoke function name for Admin.")
		//}

		/*default:
		return nil, errors.New("Invalid invoke function name for Doctor.")*/
	}

}

/*-------------------------------------------------------------------------------------------------------------*/

// returns the patient's data
// takes in 1 argument, patientID
func (t *HealthContract) getPatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 1 {
		return nil, errors.New("getPatientInfo: You need to pass 1 argument")
	}

	// get patient from table
	return nil, myTableHandler.getPatient(stub, args[0])
}

// returns the permissions of a patient
// takes in 1 argument, patientID
func (t *HealthContract) getPermissions(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 2 {
		return nil, errors.New("getPermissions: You need to pass 1 argument")
	}

	return nil, nil
}

// returns all the patient's data (array of patient structs) of a doctor
// takes in 1 argument, doctorID
func (t *HealthContract) getPatientsOfDoctor(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// check number of arguments
	if len(args) != 2 {
		return nil, errors.New("getPermissions: You need to pass 1 argument")
	}

	return nil, nil
}

// returns the doctor's data
// takes in 1 argument, doctorID
func (t *HealthContract) getDoctorInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}
