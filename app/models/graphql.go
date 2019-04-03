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
		 * curl -g 'http://localhost:9090/graphql?query={Users{Number,Name}}'
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
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Order{ID:1}{ID,Name,Total}}'
		 */
		"Order": ReadOrder,
		/*
		 * curl -g 'http://localhost:9090/graphql?query={Item{ID:101}{ID,Product{Name},Price}}'
		 */
		"Item": ReadItem,
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
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{UpdateUser(ID:1,Name:"New Tester",Password:"test"){Name}}'
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
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateOrder(Name:"Order1"){ID,Name}}'
		 */
		"CreateOrder": CreateOrder,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{UpdateOrder(ID:111,Name:"Order1",AddItem:101){ID,Name,Items{ID}}}'
		 */
		"UpdateOrder": UpdateOrder,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{CreateItem(Product:1,Price:300,Qty:2){ID,Price,Qty}}'
		 */
		"CreateItem": CreateItem,
		/*
		 *	curl -g 'http://localhost:9090/graphql?query=mutation+_{UpdateItem(ID:111,Price:"400"){ID,Product{Name},Price}}'
		 */
		"UpdateItem": UpdateItem,
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
