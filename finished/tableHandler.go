package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// config of patient table
const (
	patientTableName  = "Patients"
	iDColumn          = "ID"
	firstNameColumn   = "FirstName"
	lastNameColumn    = "LastName"
	addressColumn     = "Address"
	dateOfBirthColumn = "DateOfBirth"
)

// config of doctor table
const (
	doctorTableName = "Doctors"
	functionColumn  = "Function"
	instituteColumn = "Institute"
)

// config of permissions table
const (
	permissionsTableName = "Permissions"
	patientColumn        = "PatientID"
	doctorIDColumn       = "DoctorID"
)

type handler struct {
}

func NewTableHandler() *handler {
	return &handler{}
}

func (t *handler) createTables(stub shim.ChaincodeStubInterface) ([]byte, error) {

	// create Patient Table
	error1 := stub.CreateTable(patientTableName, []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: iDColumn, Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: firstNameColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: lastNameColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: addressColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: dateOfBirthColumn, Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if error1 != nil {
		return nil, error1
	}

	// create Doctor Table
	error2 := stub.CreateTable(doctorTableName, []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: iDColumn, Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: firstNameColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: lastNameColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: addressColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: dateOfBirthColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: functionColumn, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: instituteColumn, Type: shim.ColumnDefinition_STRING, Key: false},
	})

	if error2 != nil {
		return nil, error2
	}

	// create Permissions Table
	return nil, stub.CreateTable(permissionsTableName, []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: patientColumn, Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: doctorIDColumn, Type: shim.ColumnDefinition_STRING, Key: true},
	})

}

func (t *handler) insertPatient(stub shim.ChaincodeStubInterface, patient Patient) ([]byte, error) {

	ok, err := stub.InsertRow(patientTableName, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: patient.ID}},
			&shim.Column{Value: &shim.Column_String_{String_: patient.FirstName}},
			&shim.Column{Value: &shim.Column_String_{String_: patient.LastName}},
			&shim.Column{Value: &shim.Column_String_{String_: patient.Address}},
			&shim.Column{Value: &shim.Column_String_{String_: patient.DateOfBirth}}},
	})
	if !ok && err == nil {
		return nil, errors.New("Patient already inserted")
	}
	return nil, err
}

func (t *handler) insertDoctor(stub shim.ChaincodeStubInterface, doctor Doctor) ([]byte, error) {

	ok, err := stub.InsertRow(doctorTableName, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: doctor.ID}},
			&shim.Column{Value: &shim.Column_String_{String_: doctor.FirstName}},
			&shim.Column{Value: &shim.Column_String_{String_: doctor.LastName}},
			&shim.Column{Value: &shim.Column_String_{String_: doctor.Address}},
			&shim.Column{Value: &shim.Column_String_{String_: doctor.DateOfBirth}},
			&shim.Column{Value: &shim.Column_String_{String_: doctor.Function}},
			&shim.Column{Value: &shim.Column_String_{String_: doctor.Institute}}},
	})
	if !ok && err == nil {
		return nil, errors.New("Doctor already inserted")
	}
	return nil, err
}

func (t *handler) insertPermission(stub shim.ChaincodeStubInterface, permission Permission) ([]byte, error) {

	ok, err := stub.InsertRow(doctorTableName, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: permission.PatientID}},
			&shim.Column{Value: &shim.Column_String_{String_: permission.DoctorID}}},
	})

	if !ok && err == nil {
		return nil, errors.New("Permission already inserted")
	}

	return nil, err
}

func (t *handler) queryTable(stub shim.ChaincodeStubInterface, tablename string, columns []shim.Column) (shim.Row, error) {
	return stub.GetRow(tablename, columns)
}

func (t *handler) getPatient(stub shim.ChaincodeStubInterface, _id string) ([]string, error) {

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: _id}}
	columns = append(columns, col1)

	row, err := t.queryTable(stub, patientTableName, columns)
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return nil, errors.New("getPatient: row not found")
	}

	var pars []string
	pars = append(pars, row.Columns[0].GetString_())
	pars = append(pars, row.Columns[1].GetString_())
	pars = append(pars, row.Columns[2].GetString_())
	pars = append(pars, row.Columns[3].GetString_())
	pars = append(pars, row.Columns[4].GetString_())
	return pars, nil
}

func (t *handler) getDoctor(stub shim.ChaincodeStubInterface, _id string) ([]string, error) {

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: _id}}
	columns = append(columns, col1)

	row, err := t.queryTable(stub, doctorTableName, columns)
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return nil, errors.New("getDoctor: row not found")
	}

	var pars []string
	pars = append(pars, row.Columns[0].GetString_())
	pars = append(pars, row.Columns[1].GetString_())
	pars = append(pars, row.Columns[2].GetString_())
	pars = append(pars, row.Columns[3].GetString_())
	pars = append(pars, row.Columns[4].GetString_())
	pars = append(pars, row.Columns[5].GetString_())
	pars = append(pars, row.Columns[6].GetString_())
	return pars, nil
}

func (t *handler) getRows(stub shim.ChaincodeStubInterface, tablename string, columns []shim.Column) ([]byte, error) {

	rowChannel, err := stub.GetRows(tablename, columns)
	if err != nil {
		return nil, fmt.Errorf("getRows: %s", err)
	}

	// blob bddtests table chaincode
	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	return json.Marshal(rows)
}

func (t *handler) getDoctorPermissions(stub shim.ChaincodeStubInterface, _id string) ([]byte, error) {
	/*
		var columns []shim.Column

		col1 := shim.Column{Value: &shim.Column_String_{String_: g_id}}
		columns = append(columns, col1)

		return t.getRows(stub, permissionsTableName, columns)
	*/
	return nil, nil
}

func (t *handler) getPatientPermissions(stub shim.ChaincodeStubInterface, _id string) ([]byte, error) {
	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: _id}}
	columns = append(columns, col1)

	return t.getRows(stub, permissionsTableName, columns)
}

/*
func (t *handler) deletePatient(stub shim.ChaincodeStubInterface, _id string) error {

    err := stub.DeleteRow(
		"PatientRecords",
		[]shim.Column{shim.Column{Value: &shim.Column_String_{String_: _id}}},
	)
    if err != nil {
		return errors.New("error deleting patient")
	}
    return nil
}*/
