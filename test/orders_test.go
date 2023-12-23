package test

import (
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalOrders(t *testing.T) {
	var order = models.Order{}
	total, err := order.TotalOrders()

	assert.Equal(t, 0, total)
	assert.Nil(t, err)
}
