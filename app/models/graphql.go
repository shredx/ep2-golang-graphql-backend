package models

import "github.com/graphql-go/graphql"

//Schema is the graphql schema
var Schema graphql.Schema

//QueryConfig for graphql server
var QueryConfig = graphql.ObjectConfig{
	Name:   "Query",
	Fields: graphql.Fields{},
}

func init() {
	query := graphql.NewObject(QueryConfig)
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
	})
}
