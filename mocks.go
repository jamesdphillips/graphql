package graphql

import (
	"github.com/graphql-go/graphql"
)

//
// == Forward ==
//
// A pretty frustrating aspect of the type definitions is that they can be
// somewhat order dependent. For instance, if my Dog references Breeds and
// implements the Pet interface then I need to make sure those are loaded
// first.
//
// There are tricks to get around this (using a reference to the type and
// thunks) but the are not 100% perfect either and create a lot of needless
// code. This becomes even more tricky when generating the types.
//
// As such to get around this we generate a mock type when generating the
// object configuration; this mock only refers to the unique name of the
// component it is referencing and is replace at the time the GraphQL service
// is invoked.
//

type mockType struct{ name string }

func (o *mockType) Name() string        { return o.name }
func (o *mockType) Description() string { return "" }
func (o *mockType) String() string      { return o.name }
func (o *mockType) Error() error        { return nil }

// OutputType mocks a type (Object, Scalar, Enum, etc.)
func OutputType(name string) graphql.Output {
	return &mockType{name}
}

// InputType mocks a type (InputObject, Scalar, Enum, etc.)
func InputType(name string) graphql.Input {
	return &mockType{name}
}

// Interface mocks an interface
func Interface(name string) *graphql.Interface {
	// Unlike fields which simply require that something that implements the
	// Output interface is present object types require that references to
	// interfaces are given to config.
	//
	// Feels a bit brittle but simplest solution at this time.
	return &graphql.Interface{PrivateName: name}
}

// Object mocks an interface
func Object(name string) *graphql.Object {
	// Unlike fields which simply require that something that implements the
	// Output interface is present, schema & union types require that references
	// to object are given to config.
	//
	// Feels a bit brittle but simplest solution at this time.
	return &graphql.Object{PrivateName: name}
}

// Replace mocked types w/ instantiated counterparts
func interfacesThunk(typeMap graphql.TypeMap, cfg interface{}) interface{} {
	ints := cfg.([]*graphql.Interface)
	return graphql.InterfacesThunk(func() []*graphql.Interface {
		newInts := make([]*graphql.Interface, len(ints))
		for i, mockedInt := range ints {
			t := findType(typeMap, mockedInt.Name())
			newInt := t.(*graphql.Interface)
			newInts[i] = newInt
		}
		return newInts
	})
}

// Replace mocked types w/ instantiated counterparts
func fieldsThunk(typeMap graphql.TypeMap, fields graphql.Fields) interface{} {
	mockedFields := map[string]string{}
	for _, f := range fields {
		t := unwrapFieldType(f.Type)
		if tt, ok := t.(*mockType); ok {
			mockedFields[f.Name] = tt.Name()
		}
	}

	if len(fields) == 0 {
		return fields
	}

	return graphql.FieldsThunk(
		func() graphql.Fields {
			for fieldName, field := range fields {
				// Replace mocked instance of type
				if _, ok := mockedFields[fieldName]; ok {
					field.Type = replaceMockedType(field.Type, typeMap)
				}

				// Replace mock instances of types in arguments
				for _, arg := range field.Args {
					arg.Type = replaceMockedType(arg.Type, typeMap)
				}
			}
			return fields
		},
	)
}

// Replace mocked types w/ instantiated counterparts
func inputFieldsThunk(
	typeMap graphql.TypeMap,
	fields graphql.InputObjectConfigFieldMap,
) interface{} {
	mockedFields := []string{}
	for n, f := range fields {
		t := unwrapFieldType(f.Type)
		if _, ok := t.(*mockType); ok {
			mockedFields = append(mockedFields, n)
		}
	}

	if len(fields) == 0 {
		return fields
	}

	return graphql.InputObjectConfigFieldMapThunk(
		func() graphql.InputObjectConfigFieldMap {
			for _, name := range mockedFields {
				field := fields[name]
				field.Type = replaceMockedType(field.Type, typeMap)
			}
			return fields
		},
	)
}

func replaceMockedType(t graphql.Type, m graphql.TypeMap) graphql.Type {
	switch tt := t.(type) {
	case *graphql.NonNull:
		tt.OfType = replaceMockedType(tt.OfType, m)
		return tt
	case *graphql.List:
		tt.OfType = replaceMockedType(tt.OfType, m)
		return tt
	case *mockType:
		return findType(m, t.Name())
	default:
		return t
	}
}

func unwrapFieldType(t graphql.Type) graphql.Type {
	if tt, ok := t.(*graphql.NonNull); ok {
		t = unwrapFieldType(tt.OfType)
	} else if tt, ok := t.(*graphql.List); ok {
		t = unwrapFieldType(tt.OfType)
	}
	return t
}
