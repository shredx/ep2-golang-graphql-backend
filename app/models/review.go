package models

import "github.com/graphql-go/graphql"

/*
 * This file contains the model of review
 */

//Review given by a user for the product
type Review struct {
	//ID of the review
	ID uint `gorm:"primary_key"`
	//Review of the user
	Review string
	//Rating given by the user
	Rating int
}

//ReviewConfig is the object config for review interface
var ReviewConfig = graphql.ObjectConfig{
	Name: "Review",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Review": &graphql.Field{
			Type: graphql.String,
		},
		"Rating": &graphql.Field{
			Type: graphql.Int,
		},
	},
}

//ReviewSchema is the schema of the review model
var ReviewSchema = graphql.NewObject(ReviewConfig)

//ReadReview will read a review
var ReadReview = &graphql.Field{
	Type:        ReviewSchema, // the return type for this field
	Description: "Get a single review",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, ok := params.Args["ID"]
		review := Review{}

		if !ok {
			//id is not given
			return review, nil
		}

		//finding the review from the db
		DB.First(&review, uint(id.(int)))

		// return the new User object that we supposedly save to DB
		return review, nil
	},
}
