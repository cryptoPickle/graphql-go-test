package gql

import (
	"github.com/graphql-go/graphql"
	postgress2 "testgraphql/pkg/postgress"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *postgress2.Db) *Root{
	resolver := Resolver{db: db}

	root := Root{Query:graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type:              graphql.NewList(User),
					Args:              graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type:         graphql.String,
						},
					},
					Resolve: resolver.UserResolver,
				},
			},
		},
		)}

	return &root
}
