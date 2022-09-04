package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testDSN DSN = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
)

func TestConnectRepository(t *testing.T) {
	repository, err := NewGormRepository(testDSN)
	assert.Nil(t, err)
	assert.NotNil(t, repository)
}

func TestMigrate(t *testing.T) {
	repository, err := NewGormRepository(testDSN)
	assert.Nil(t, err)
	assert.NotNil(t, repository)
	err = repository.Migrate()
	assert.Nil(t, err)
}
