package test

import (
	"fmt"
	"io"
	"log"
	"logistica/app/models"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
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

	fmt.Println(data)

	assert.Nil(t, err)
	// assert.Equal(t, "muhammad dhimas arista", data[1]["name"])
}

func TestGetEmployeeLastId(t *testing.T) {
	var data = models.Employee{}
	lastId, err := data.LastId()
	if err != nil {
		log.Println(err)
	}

	assert.Nil(t, err)
	fmt.Println(lastId)
}

func TestCheckId(t *testing.T) {
	app := fiber.New()

	request := httptest.NewRequest("GET", "localhost:5500/employee/check/100011", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	// assert.Contains(t, bytes, )

	fmt.Println(string(bytes))
}

func TestTotal(t *testing.T) {
	var employee = models.Employee{}
	total, err := employee.Count()
	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 2, total)
}

func TestNewEmployee(t *testing.T) {
	var employee = models.Employee{}
	lastId, _ := employee.LastId()

	_, err := employee.NewEmployee(lastId+1, "Test New Employee", "Unknown", "+628000000000", 2224)
	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
}
