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
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// getCallerData is used to get the ecert and attributes of the caller.
func (t *HealthContract) getCallerData(stub shim.ChaincodeStubInterface) (string, string, error) {

	username, err := stub.ReadCertAttribute("userName")
	if err != nil {
		fmt.Println("userNname" + err.Error())
	}

	fmt.Println("UserName: " + string(username))

	typeOfUser, err := stub.ReadCertAttribute("typeOfUser")
	if err != nil {
		fmt.Println("typeOfUser" + err.Error())
	}
	fmt.Println("TypeOfUser: " + string(typeOfUser))

	return string(username), string(typeOfUser), nil

}
