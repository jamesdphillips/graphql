package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockType(t *testing.T) {
	mock := &mockType{"test"}
	assert.Equal(t, mock.Name(), "test")
	assert.Equal(t, mock.Description(), "")
	assert.Equal(t, mock.String(), "test")
	assert.NoError(t, mock.Error())
}
