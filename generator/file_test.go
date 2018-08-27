package generator

import (
	"bytes"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/require"
)

type testSaver struct {
	out string
}

func (t *testSaver) Save(_ string, f *jen.File) error {
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return err
	}
	t.out = buf.String()
	return nil
}

func TestLoadDir(t *testing.T) {
	fs, err := ParseDir("./kitchen-sink-schema")
	require.NoError(t, err)
	require.NotEmpty(t, fs)
	require.NoError(t, fs.Validate())
}
