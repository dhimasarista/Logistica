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
