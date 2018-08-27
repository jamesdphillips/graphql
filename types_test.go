package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewType(t *testing.T) {
	myType := NewType("type", InterfaceKind)
	assert.Equal(t, myType.Name(), "type")
	assert.Equal(t, myType.Kind(), InterfaceKind)
}
