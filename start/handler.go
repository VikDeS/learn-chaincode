package main

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type handler struct {
}

// config of chaincode table
const (
	tableColumn             = "PatientRecords"
	columnAccountID         = "Account"
	columnSocialSecurityNr  = "Ssnr"
	columnDateOfBirth       = "DateOfBirth"
    columnAddress           = "Address"
)

func NewHandler() *handler {
	return &handler{}
}

func (t *handler) createTable(stub shim.ChaincodeStubInterface) error {

	return stub.CreateTable(tableColumn, []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: columnAccountID, Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: columnSocialSecurityNr, Type: shim.ColumnDefinition_UINT64, Key: false},
		&shim.ColumnDefinition{Name: columnDateOfBirth, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnAddress, Type: shim.ColumnDefinition_STRING, Key: false},
	})

}

func (t *handler) insertPatient(stub shim.ChaincodeStubInterface, 
    _id string, 
    _ssnr uint64, 
    _birth string, 
    _address string) error {

    ok, err := stub.InsertRow(tableColumn, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: _id}},
            &shim.Column{Value: &shim.Column_Uint64{Uint64: _ssnr}},
			&shim.Column{Value: &shim.Column_String_{String_: _birth}},
			&shim.Column{Value: &shim.Column_String_{String_: _address}}},
	})
    if !ok && err == nil {
		return errors.New("Patient already inserted")
	}
    return nil
}

func (t *handler) deletePatient(stub shim.ChaincodeStubInterface, _id string) error {

    err := stub.DeleteRow(
		"PatientRecords",
		[]shim.Column{shim.Column{Value: &shim.Column_String_{String_: _id}}},
	)
    if err != nil {
		return errors.New("error deleting patient")
	}
    return nil
}

func (t *handler) queryPatient(stub shim.ChaincodeStubInterface, _id string) (string, uint64, string, string, error) {

    row, err := t.queryTable(stub, _id)
	if err != nil {
		return "", 0, "", "", err
	}
	if len(row.Columns) == 0 {
		return "", 0, "", "",  errors.New("row not found")
	}

	return row.Columns[0].GetString_(), row.Columns[1].GetUint64(), row.Columns[2].GetString_(), row.Columns[3].GetString_(), nil
}

func (t *handler) queryTable(stub shim.ChaincodeStubInterface, _id string) (shim.Row, error) {

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: _id}}
	columns = append(columns, col1)

	return stub.GetRow(tableColumn, columns)
}