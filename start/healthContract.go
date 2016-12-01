/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var patientRecords []Patient
var myhandler = NewHandler()

type Patient struct {
	ssnr int
	name string
	//gender Gender
	birthdate string
	address   string
	//permissionedViewers []int
	//medicalData []byte
}

type HealthContract struct {
}

/*
type Gender struct {}

const (
    MALE Gender = 1 + iota
    FEMALE
)*/
//Init the chaincode asigned the value "0" to the counter in the state.
func (t *HealthContract) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//err := stub.PutState("counter", []byte("0"))
	return nil, nil
}

//Invoke Transaction makes increment counter
func (t *HealthContract) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case "addPatient":
		return t.addPatient(stub, args)
	case "changePatientInfo":
		return t.changePatientInfo(stub, args)
	case "removePatient":
		return t.removePatient(stub, args)
	case "givePermission":
		return t.givePermission(stub, args)
	case "removePermission":
		return t.removePermission(stub, args)
	default:
		return nil, errors.New("Invalid invoke function name.")
	}

	/*if function == "increment" {

		return t.increment(stub, args)
	}
	if function == "decrement" {
		counter, err := stub.GetState("counter")
		if err != nil {
			return nil, err
		}
		var cInt int
		cInt, err = strconv.Atoi(string(counter))
		if err != nil {
			return nil, err
		}
		cInt = cInt - 1
		counter = []byte(strconv.Itoa(cInt))
		stub.PutState("counter", counter)
	}*/
	return nil, nil

}

func (t *HealthContract) addPatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// check if right nr of arguments
	ssnr, err := strconv.ParseUint(args[1], 10, 64)

	if err != nil {
		return nil, err
	}
	// addPatient to
	//append(patientRecords , len(Patient))
	/*
		    counter, err := stub.GetState("counter")
			if err != nil {
				return nil, err
			}
			var cInt int
			cInt, err = strconv.Atoi(string(counter))
			if err != nil {
				return nil, err
			}
			cInt = cInt + 1
			counter = []byte(strconv.Itoa(cInt))
			stub.PutState("counter", counter)*/

	return nil, myhandler.insertPatient(stub, args[0], ssnr, args[2], args[3])

}

func (t *HealthContract) changePatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// check if right nr of arguments

	// get first argument is object with all the new info
	var err = stub.PutState("logger", []byte("change patientInfo of "+args[0]))
	return nil, err
}

func (t *HealthContract) removePatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err = stub.PutState("logger", []byte("remove patient "+args[0]))
	return nil, err
}

func (t *HealthContract) givePermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err = stub.PutState("logger", []byte("give permission to "+args[0]))
	return nil, err
}

func (t *HealthContract) removePermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err = stub.PutState("logger", []byte("remove permission of "+args[0]))
	return nil, err
}

// Query callback representing the query of a chaincode
func (t *HealthContract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	/*var logger
	logger, err = stub.GetState("logger")
	*/
	switch function {
	case "getPatientInfo":
		t.getPatientInfo(stub, args)
	case "getPatientData":
		t.getPatientData(stub, args)
	case "getPermissions":
		t.getPermissions(stub, args)
	default:
		return nil, errors.New("Invalid query function name.")
	}
	return nil, nil

	/*
		var err error

		// Get the state from the ledger
		Avalbytes, err := stub.GetState("counter")
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for counter\"}"
			return nil, errors.New(jsonResp)
		}

		if Avalbytes == nil {
			jsonResp := "{\"Error\":\"Nil amount for counter\"}"
			return nil, errors.New(jsonResp)
		}

		jsonResp := "{\"Name\":\"counter\",\"Amount\":\"" + string(Avalbytes) + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return Avalbytes, nil*/
}

func (t *HealthContract) getPatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *HealthContract) getPatientData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *HealthContract) getPermissions(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func main() {
	err := shim.Start(new(HealthContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
