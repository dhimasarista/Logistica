package test

import (
	"logistica/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLDB(t *testing.T) {
	var db = config.ConnectSQLDB()
	err := db.Ping()

	assert.Nil(t, err)
}
