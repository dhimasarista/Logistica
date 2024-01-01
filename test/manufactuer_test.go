package test

import (
	"fmt"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManufacturerFindAll(t *testing.T) {
	manufacturer := models.Manufacturer{}

	manufacturers, err := manufacturer.FindAll()
	fmt.Println(manufacturers)

	assert.Nil(t, err)
}

func TestManufacturerLastId(t *testing.T) {
	manufacturer := models.Manufacturer{}
	lastId, err := manufacturer.LastId()
	fmt.Println(lastId)

	assert.Nil(t, err)
}
func TestNewManufacturer(t *testing.T) {
	manufacturer := models.Manufacturer{}
	lastId, _ := manufacturer.LastId()

	err := manufacturer.NewManufacturer(lastId+1, "Hello")
	assert.Nil(t, err)
}
