package graphql

import (
	graphql1 "github.com/graphql-go/graphql"
	ast "github.com/graphql-go/graphql/language/ast"
	mapstructure "github.com/mitchellh/mapstructure"
)

func registerSchema(svc *Service) {
	svc.RegisterSchema(_SchemaDesc)
}
func _SchemaConfigFn() graphql1.SchemaConfig {
	return graphql1.SchemaConfig{Query: Object("Foo")}
}

var _SchemaDesc = SchemaDesc{Config: _SchemaConfigFn}

type fooOneFieldResolverArgs struct {
	First int
}

type fooOneFieldResolverParams struct {
	ResolveParams
	Args fooOneFieldResolverArgs
}

type fooOneFieldResolver interface {
	One(p fooOneFieldResolverParams) (float64, error)
}

type fooFieldResolvers interface {
	fooOneFieldResolver
}

type fooAliases struct{}

func (_ fooAliases) One(p fooOneFieldResolverParams) (float64, error) {
	return 1.0, nil
}

func registerFoo(svc *Service, impl fooFieldResolvers) {
	svc.RegisterObject(_ObjectTypeFooDesc, impl)
}
func _ObjTypeFooOneHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(fooOneFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := fooOneFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.One(frp)
	}
}

func _ObjectTypeFooConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "self descriptive",
		Fields: graphql1.Fields{"one": &graphql1.Field{
			Args: graphql1.FieldConfigArgument{"first": &graphql1.ArgumentConfig{
				Description: "self descriptive",
				Type:        graphql1.Int,
			}},
			DeprecationReason: "",
			Description:       "self descriptive",
			Name:              "one",
			Type:              graphql1.NewNonNull(graphql1.Float),
		}},
		Interfaces: []*graphql1.Interface{
			Interface("Bar"),
		},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			panic("Unimplemented; see FooFieldResolvers.")
		},
		Name: "Foo",
	}
}

var _ObjectTypeFooDesc = ObjectDesc{
	Config:        _ObjectTypeFooConfigFn,
	FieldHandlers: map[string]FieldHandler{"one": _ObjTypeFooOneHandler},
}

type bazTwoFieldResolver interface {
	Two(p ResolveParams) (interface{}, error)
}

type bazFieldResolvers interface {
	bazTwoFieldResolver
}

type bazAliases struct{}

func (_ bazAliases) Two(p ResolveParams) (interface{}, error) {
	val, err := DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

func registerBaz(svc *Service, impl bazFieldResolvers) {
	svc.RegisterObject(_ObjectTypeBazDesc, impl)
}
func _ObjTypeBazTwoHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(bazTwoFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Two(frp)
	}
}

func _ObjectTypeBazConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "self descriptive",
		Fields: graphql1.Fields{"two": &graphql1.Field{
			Args:              graphql1.FieldConfigArgument{},
			DeprecationReason: "",
			Description:       "self descriptive",
			Name:              "two",
			Type:              graphql1.NewNonNull(OutputType("Foo")),
		}},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			panic("Unimplemented; see BazFieldResolvers.")
		},
		Name: "Baz",
	}
}

var _ObjectTypeBazDesc = ObjectDesc{
	Config:        _ObjectTypeBazConfigFn,
	FieldHandlers: map[string]FieldHandler{"two": _ObjTypeBazTwoHandler},
}

func registerUrl(svc *Service, impl ScalarResolver) {
	svc.RegisterScalar(_ScalarTypeUrlDesc, impl)
}

var _ScalarTypeUrlDesc = ScalarDesc{Config: func() graphql1.ScalarConfig {
	return graphql1.ScalarConfig{
		Description: "self descriptive",
		Name:        "Url",
		ParseLiteral: func(_ ast.Value) interface{} {

			panic("Unimplemented; see ScalarResolver.")
		},
		ParseValue: func(_ interface{}) interface{} {

			panic("Unimplemented; see ScalarResolver.")
		},
		Serialize: func(_ interface{}) interface{} {

			panic("Unimplemented; see ScalarResolver.")
		},
	}
}}

type urlImpl struct{}

func (urlImpl) ParseLiteral(v ast.Value) interface{} {
	panic("ParseLiteral called")
}

func (urlImpl) ParseValue(v interface{}) interface{} {
	panic("ParseValue called")
}

func (urlImpl) Serialize(v interface{}) interface{} {
	panic("Serialize called")
}

func registerBar(svc *Service, impl InterfaceTypeResolver) {
	svc.RegisterInterface(_InterfaceTypeBarDesc, impl)
}
func _InterfaceTypeBarConfigFn() graphql1.InterfaceConfig {
	return graphql1.InterfaceConfig{
		Description: "self descriptive",
		Fields: graphql1.Fields{"one": &graphql1.Field{
			Args: graphql1.FieldConfigArgument{"first": &graphql1.ArgumentConfig{
				Description: "self descriptive",
				Type:        graphql1.Int,
			}},
			DeprecationReason: "",
			Description:       "self descriptive",
			Name:              "one",
			Type:              graphql1.NewNonNull(graphql1.Float),
		}},
		Name: "Bar",
		ResolveType: func(_ graphql1.ResolveTypeParams) *graphql1.Object {

			panic("Unimplemented; see InterfaceTypeResolver.")
		},
	}
}

var _InterfaceTypeBarDesc = InterfaceDesc{Config: _InterfaceTypeBarConfigFn}

func registerFooBar(svc *Service, impl UnionTypeResolver) {
	svc.RegisterUnion(_UnionTypeFooBarDesc, impl)
}
func _UnionTypeFooBarConfigFn() graphql1.UnionConfig {
	return graphql1.UnionConfig{
		Description: "self descriptive",
		Name:        "FooBar",
		ResolveType: func(_ graphql1.ResolveTypeParams) *graphql1.Object {

			panic("Unimplemented; see UnionTypeResolver.")
		},
		Types: []*graphql1.Object{
			Object("Foo"),
			Object("Baz")},
	}
}

var _UnionTypeFooBarDesc = UnionDesc{Config: _UnionTypeFooBarConfigFn}

func registerEnum(svc *Service) {
	svc.RegisterEnum(_EnumTypeEnumDesc)
}
func _EnumTypeEnumConfigFn() graphql1.EnumConfig {
	return graphql1.EnumConfig{
		Description: "self descriptive",
		Name:        "Enum",
		Values: graphql1.EnumValueConfigMap{"VALUE": &graphql1.EnumValueConfig{
			DeprecationReason: "",
			Description:       "self descriptive",
			Value:             "VALUE",
		}},
	}
}

var _EnumTypeEnumDesc = EnumDesc{Config: _EnumTypeEnumConfigFn}

func registerInput(svc *Service) {
	svc.RegisterInput(_InputTypeInputDesc)
}
func _InputTypeInputConfigFn() graphql1.InputObjectConfig {
	return graphql1.InputObjectConfig{
		Description: "self descriptive",
		Fields: graphql1.InputObjectConfigFieldMap{"three": &graphql1.InputObjectFieldConfig{
			Description: "self descriptive",
			Type:        graphql1.NewNonNull(graphql1.String),
		}},
		Name: "Input",
	}
}

var _InputTypeInputDesc = InputDesc{Config: _InputTypeInputConfigFn}
