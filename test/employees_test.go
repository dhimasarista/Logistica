package test

import (
	"fmt"
	"log"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployeeById(t *testing.T) {
	var data = models.Employee{}
	err := data.GetById(100011)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 100011, int(data.ID.Int64))
}

func TestFindAllEmployee(t *testing.T) {
	var employees = models.Employee{}
	data, err := employees.FindAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(data[1]["name"])
}
