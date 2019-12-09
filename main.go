package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"testgraphql/pkg/gql"
	"testgraphql/pkg/server"

	"github.com/go-chi/chi"

	postgress2 "testgraphql/pkg/postgress"
)

func initialiseApi()(*chi.Mux, *postgress2.Db) {
	router := chi.NewRouter()

	db, err := postgress2.New(postgress2.ConnString("localhost", "user", "dbName", 5432))

	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(graphql.SchemaConfig{Query:rootQuery.Query})

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{GqlSchema:&sc}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
		)

	router.Post("/graphql", s.GraphQL())

	return router, db
}


func main(){
	router, db := initialiseApi()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}