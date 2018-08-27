package graphql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	svc := NewService()
	require.NotNil(t, svc)
}

func TestRegerate(t *testing.T) {
	svc := NewService()
	registerTestTypes(svc)
	err := svc.Regenerate()
	require.NoError(t, err)
}

func TestDo(t *testing.T) {
	svc := NewService()
	registerTestTypes(svc)
	svc.Regenerate()

	ctx := context.Background()
	res := svc.Do(ctx, "query { one(first: [{three: \"four\"}]) }", map[string]interface{}{})
	require.Empty(t, res.Errors)
	assert.NotEmpty(t, res.Data)
}

func registerTestTypes(svc *Service) {
	registerFoo(svc, &fooAliases{})
	registerBaz(svc, &bazAliases{})
	registerUrl(svc, urlImpl{})
	registerBar(svc, nil)
	registerFooBar(svc, nil)
	registerEnum(svc)
	registerInput(svc)
	registerSchema(svc)
}
