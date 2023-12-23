package test

import (
	"fmt"
	"logistica/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumEarnings(t *testing.T) {
	var earning = models.Earning{}
	sum, err := earning.TotalEarnings()

	fmt.Println(sum)
	assert.Nil(t, err)
}
