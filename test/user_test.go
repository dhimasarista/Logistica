package test

import (
	"fmt"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	var user = &models.User{}
	err := user.GetByID(2)
	fmt.Println(user.ID)
	assert.Nil(t, err)
}
