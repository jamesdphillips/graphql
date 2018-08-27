<p align="center">
  <a href="https://www.graphql.org/">
    <img alt="senu" src="https://graphql.org/img/logo.svg" width="144">
  </a>
</p>

<h3 align="center">
  GraphQL
</h3>

<p align="center">
  Go Flavoured GraphQL module for implementing fast, type safe
  <a href="https://graphql.org">GraphQL</a> services. Extracted from
  <a href="https://github.com/sensu/sensu-go">sensu/sensu-go</a>. Represents
  code in production by organizations running
  <a href="https://sensu.io">Sensu</a> version 2.
</p>

<p align="center">
  <a href="https://circleci.com/gh/sensu/sensu-go/tree/master"><img src="https://circleci.com/gh/sensu/sensu-go/tree/master.svg?style=svg"></a>
</p>

## Suggested Reading

- [Introduction to GraphQL](https://www.graphql.org/learn)
- [GraphQL IDL](https://www.graphql.org/learn)
- [GraphQL Mutations](http://graphql.org/learn/queries/#mutations)

## Basic Usage

For those familiar with implementing a gRPC service in golang should feel right
at home.

1.  Add new type / field definition(s).

    ```graphql
    # ./backend/apid/graphql/schema/dog.graphql

    """
    A Dog are the best pets.
    """
    type Dog implements Named {
      name: String!
      profilePicture(size: Size): String
      friends: [Pet]!
    }

    interface Named {
      name: String!
    }

    type Rabbit implements Named {
      name: String!
      friends: [Pet]!
      isFluffy: Boolean!
    }

    input Size {
      width: Int
      height: Int
      density: float
    }

    union Pet = Dog | Rabbit
    ```

2.  Next you'll want to generate the Go code for these new types. To do this we
    use the `gengraphql` script.

    Run:

    ```shell
    graphqlgen ./path-to-my-schema
    ```

3.  Next we need to tell our service how the types themselves are implemented.
    To give an example, when a user selects a dog's friends, the service needs
    to know _how_ to retrieve those details so that it can be display them to
    the user.

    An example implementation:

    ```go
    // backend/apid/graphql/dog.go

    type dogFieldResolvers struct {
      // the autogenerated "aliases" use reflection under the hood to implement most
      // most the field resolvers; this allows us to write a bit less code with a
      // small runtime cost.
      *schema.DogAliases

      controller FriendsController
      logger     logrus.Entry
    }

    // Use controller to retrieve our dog's friends.
    func (fr *dogFieldResolvers) Friends(p graphql.ResolveParams) (interface{}, error) {
      dog := p.Source.(*types.Dog)
      ctx := p.Context
      friends, err := fr.controller.ListFriendos(ctx, dog)
      return friends, err
    }

    // IsTypeOf is used to determine if a given value is associated with the Dog type
    func (fr *dogFieldResolvers) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
      _, ok := s.(*types.Dog)
      return ok
    }
    ```

4.  Finally we need to register the new type(s) and any of the implementation
    details with our service.

    ```go
    // backend/apid/graphql/service.go

    fund NewService(c Config) *Service {
      // ...
      dogImpl := dogFieldResolvers{controller: ..., logger: ...}
      // ...

      // ...
      schema.RegisterDog(svc, dogImpl) // include fieldresolvers.
      schema.RegisterSize(svc)         // unlike object type's inputs do not require any additonal implemtation details.
      // ...

      // configures registered types and implementations so that service is ready to
      // accept queries.
      service.Reconfigure()
      return service
    }
    ```

## Mutations

Mutations are how the client modifies the server-side data.

-  [Reference](http://graphql.org/learn/queries/#mutations)
-   Sensu follow's Relay's [mutation conventions](https://facebook.github.io/relay/docs/en/graphql-server-specification.html#mutations). Each mutation should consist of three
    elements. An input object type that describes the parameters to the mutation,
    an object type that describes the return values, and finally a field on the
    `Mutation` to be used as the entry point.

    ```graphql
    # mutations.graphql
    type Mutation {
      # ...
      addRole(inputs: AddRoleInput) AddRolePayload
      # ...
    }

    input AddRoleInput {
      # Used by a client to keep track of in-flight mutations
      clientMutationId: String!
      userId: ID!
      roleId: ID!
    }

    type AddRolePayload {
      clientMutationId: String!
      user: User!
      role: Role!
    }
    ```

## Deprecation

- Fields should not be removed until we can be confident that no clients are
  using the field.
- GraphQL supports @deprecated directive for marking a field as deprecated.

    ```graphql
    type MyType {
      one: String! @deprecated
      two: String! @deprecated(reason: "Two is bad number.")
      three: String!
    }
    ```

## File Conventions

- **Type Definitions** live in the `schema` package, and use the file extension
  `.graphql`.
  - When the type(s) they match an internal type defined in the
  `types` package the filenames should ideally match. (Eg. `entity.go`
  `entity.graphql`.)
  - Ideally all types and fields are
- **FieldResolvers** live in the `graphql` package.
  - Filenames should match the same name of the graphql file it is implementing.

[Sensu]:https://www.sensu.io
[sensu/sensu-go]:https://github.com/sensu/sensu-go
