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

func TestProductStock(t *testing.T) {
	product := models.Product{}

	stocks, err := product.UpdateStocks(1023, 28)
	fmt.Println(stocks.RowsAffected())

	assert.Nil(t, err)
}
func TestLastStock(t *testing.T) {
	product := models.Product{}

	stocks, err := product.LastStocks(1023)
	fmt.Println(stocks)

	assert.Nil(t, err)
}

func TestLastIdProduct(t *testing.T) {
	product := models.Product{}

	lastId, err := product.LastId()

	assert.Nil(t, err)
	assert.Equal(t, 1024, lastId)
}
