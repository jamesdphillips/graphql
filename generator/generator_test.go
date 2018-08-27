package generator

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKitchenSinkExample(t *testing.T) {
	file, err := ParseFile("./kitchen-sink-schema/schema.graphql")
	require.NoError(t, err)
	require.NotNil(t, file)
	require.NoError(t, file.Validate())

	files := GraphQLFiles{file}
	generator := New(files)
	require.NotNil(t, generator)

	saver := testSaver{}
	generator.Saver = &saver

	gerr := generator.Run()
	require.NoError(t, gerr)
	assert.NotEmpty(t, saver.out)

	perr := parseSrc(saver.out)
	assert.NoError(t, perr)
}

func parseSrc(src string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	return err
}
