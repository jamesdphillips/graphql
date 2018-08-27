package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultResolver(t *testing.T) {
	animal := struct {
		Name string `json:"firstname"`
	}{"bob"}
	testCases := []struct {
		desc   string
		source interface{}
		field  string
		out    interface{}
	}{
		{
			desc:   "field on struct",
			source: animal,
			field:  "name",
			out:    "bob",
		},
		{
			desc:   "field on struct w/ tag",
			source: animal,
			field:  "firstname",
			out:    "bob",
		},
		{
			desc:   "missing field on struct",
			source: animal,
			field:  "surname",
			out:    nil,
		},
		{
			desc:   "field on map",
			source: map[string]interface{}{"name": "bob"},
			field:  "name",
			out:    "bob",
		},
		{
			desc:   "missing field on map",
			source: map[string]interface{}{"name": "bob"},
			field:  "firstname",
			out:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := DefaultResolver(tc.source, tc.field)
			require.NoError(t, err)
			assert.EqualValues(t, result, tc.out)
		})
	}
}
