package test

import (
	"fmt"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderDetail(t *testing.T) {
	var od = models.OrderDetail{}
	res, err := od.NewOrder("anto", "082838248", "jl. kenangan")

	assert.Nil(t, err)
	fmt.Println(res)
}
func TestGetOrderDetail(t *testing.T) {
	var od = &models.OrderDetail{}
	err := od.GetByID(2)
	assert.Nil(t, err)
	assert.Equal(t, "anto", od.Buyer.String)
}
