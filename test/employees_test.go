package test

import (
	"database/sql"
	"fmt"
	"log"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sukses
func TestGetEmployeeById(t *testing.T) {
	var data = models.Employee{}
	err := data.GetById(1)
	if err != nil {
		log.Println(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 1, int(data.ID.Int64))
	assert.Equal(t, "administrator", data.Name.String)
}

// Sukses
func TestFindAllEmployee(t *testing.T) {
	var employees = models.Employee{}
	_, err := employees.FindAll()
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
}

// Sukses
func TestGetEmployeeLastId(t *testing.T) {
	var data = models.Employee{}
	lastId, err := data.LastId()
	if err != nil {
		log.Println(err)
	}

	assert.Nil(t, err)
	fmt.Println(lastId)
}

// Sukses
func TestTotal(t *testing.T) {
	var employee = models.Employee{}
	total, err := employee.Count()
	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 1, total)
}

// Sukses
func TestNewEmployee(t *testing.T) {
	var employee = models.Employee{
		ID:          sql.NullInt64{Int64: 0},
		Name:        sql.NullString{String: "Unknown"},
		Address:     sql.NullString{String: "Unknown"},
		NumberPhone: sql.NullString{String: "Unknown"},
		PositionID:  sql.NullInt64{Int64: 2222},
	}

	err := employee.NewEmployee()
	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
}

// Sukses
func TestUpdateEmployee(t *testing.T) {
	var employee = models.Employee{
		ID:          sql.NullInt64{Int64: 100021},
		Name:        sql.NullString{String: "Unknown"},
		Address:     sql.NullString{String: "Unknown"},
		NumberPhone: sql.NullString{String: "Unknown"},
		PositionID:  sql.NullInt64{Int64: 2222},
	}
	err := employee.UpdateEmployee()
	assert.Nil(t, err)
}
