package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID("test", "123")
	assert.Equal(t, ID("test_123"), id)
}
