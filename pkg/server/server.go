package server

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"net/http"
	"testgraphql/pkg/gql"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", http.StatusBadRequest )
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)

		if err != nil {
			http.Error(w, "Error parsing JSON request Body", http.StatusBadRequest)
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)

		render.JSON(w, r, result)
	}
}
