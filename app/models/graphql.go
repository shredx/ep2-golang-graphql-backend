package models

import (
	"github.com/graphql-go/graphql"
)

//Schema is the graphql schema
var Schema graphql.Schema

//QueryConfig for graphql server query
var QueryConfig = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/*
		 * curl -g 'http://localhost:9090/graphql?query={User(ID:1){Number,Name}}'
		 */
		"User": ReadUser,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={User{Number,Name}}'
		 */
		"Users": ReadUsers,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Review(ID:1){Review,Rating}}'
		 */
		"Review": ReadReview,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Tag(ID:1){Name}}'
		 */
		"Tag": ReadTag,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Tags{ID,Name}}'
		 */
		"Tags": ReadTags,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Product(ID:1){ID,Name,Brand}}'
		 */
		"Product": ReadProduct,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Category(ID:1){ID,Name,Description}}'
		 */
		"Category": ReadCategory,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Categories{ID,Name,Description}}'
		 */
		"Categories": ReadCategories,
	},
}

//MutationConfig for the graphql server mutation
var MutationConfig = graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateUser(Number:"9999990000",Name:"Tester",Password:"test"){ID}}'
		 */
		"CreateUser": CreateUser,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateUser(ID:1,Number:"9999990000",Name:"New Tester",Password:"test"){Name}}'
		 */
		"UpdateUser": UpdateUser,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateTag(Name:"Electronics"){ID,Name}}'
		 */
		"CreateTag": CreateTag,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{UpdateTag(ID:1,Name:"Laptops"){ID,Name}}'
		 */
		"UpdateTag": UpdateTag,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateProduct(Name:"Lenovo",Price:40000){ID,Name}}'
		 */
		"CreateProduct": CreateProduct,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{UpdateProduct(ID:1,Name:"Lenovo%20Ideapad"){ID,Name}}'
		 */
		"UpdateProduct": UpdateProduct,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{ProductReview(ID:1,Review:"Good%20Product",Rating:4,Add:true){ID,Name,Reviews{Review,Rating}}}'
		 */
		"ProductReview": ProductReview,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateCategory(Name:"Electronics"){ID,Name}}'
		 */
		"CreateCategory": CreateCategory,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{UpdateCategory(ID:1,Name:"Computers%20Gadgets"){ID,Name}}'
		 */
		"UpdateCategory": UpdateCategory,
	},
}

func init() {
	/*
	 * Run the model initializations
	 * Set the query and mutations
	 */
	//Model initializations
	InitProduct()
	InitCategories()

	//Set the query and mutations
	query := graphql.NewObject(QueryConfig)
	mutation := graphql.NewObject(MutationConfig)
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}
