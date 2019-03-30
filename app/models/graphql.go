package models

import "github.com/graphql-go/graphql"

//Schema is the graphql schema
var Schema graphql.Schema

//QueryConfig for graphql server query
var QueryConfig = graphql.ObjectConfig{
	Name:   "Query",
	Fields: graphql.Fields{},
}

//MutationCOnfig for the graphql server mutation
var MutationConfig = graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: graphql.Fields{},
}

func init() {
	query := graphql.NewObject(QueryConfig)
	mutation := graphql.NewObject(MutationConfig)
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}
