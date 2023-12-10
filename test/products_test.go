package test

import (
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductByID(t *testing.T) {
	product := &models.Product{}
	err := product.GetById(1023)

	assert.Nil(t, err)
}
