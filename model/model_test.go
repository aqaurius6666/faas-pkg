package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	m, err := NewModel(Option{
		Repository: nil,
	})
	assert.Nil(t, err)
	assert.NotNil(t, m)
}
