package test

import (
	"fmt"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductByID(t *testing.T) {
	product := &models.Product{}
	err := product.GetById(1023)

	assert.Nil(t, err)
}

func TestProductFindAll(t *testing.T) {
	product := &models.Product{}
	data, err := product.FindAll()

	fmt.Println(data)

	assert.Nil(t, err)
}

func TestCountProducts(t *testing.T) {
	product := models.Product{}

	data, err := product.Count()

	fmt.Println(data)

	assert.Nil(t, err)
}
