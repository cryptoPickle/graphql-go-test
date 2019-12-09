package gql

import (
	"github.com/graphql-go/graphql"

	"testgraphql/pkg/postgress"
)

type Resolver struct {
	db *postgress.Db
}


func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)

	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}

	return nil, nil
};