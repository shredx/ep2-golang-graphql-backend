package models

import (
	"github.com/graphql-go/graphql"
)

/*
 * This file contains the model defintions for Product
 */

//Product is the data structure of the product
type Product struct {
	//ID of the product
	ID          uint   `gorm:"primary_key"`
	Name        string //Name of the product
	Brand       string //Brand to which the product belongs to
	Description string //Description of the product
	//Tags is the tags associated with the product
	Tags  []Tag   `gorm:"many2many:product_tags;"`
	Price float64 //Price of the product
	Image string  //Image url of the product
	//Reviews given to this product
	Reviews []Review `gorm:"many2many:product_reviews;"`
	ItemID  uint     // foreign_key for Item.Product
}

//ProductConfig is the config for the product object
var ProductConfig = graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"Brand": &graphql.Field{
			Type: graphql.String,
		},
		"Description": &graphql.Field{
			Type: graphql.String,
		},
		"Tags": &graphql.Field{
			Type: graphql.NewList(TagSchema),
		},
		"Price": &graphql.Field{
			Type: graphql.Float,
		},
		"Image": &graphql.Field{
			Type: graphql.String,
		},
		"Reviews": &graphql.Field{
			Type: graphql.NewList(ReviewSchema),
		},
	},
}

//ProductSchema is the schema of the product model
var ProductSchema = graphql.NewObject(ProductConfig)

//ReadProduct will read a Product
var ReadProduct = &graphql.Field{
	Type:        ProductSchema, // the return type for this field
	Description: "Get a single Product",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		Product := Product{ID: uint(id)}

		//finding the Product from the db
		DB.First(&Product).Related(&Product.Tags, "Tags").Related(&Product.Reviews, "Reviews")

		// return the new Product object that we supposedly save to DB
		return Product, nil
	},
}

//ProductArgumentConfig is argument config required for the product
var ProductArgumentConfig = graphql.FieldConfigArgument{
	"Name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"Brand": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"Description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"Tags": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.Int),
	},
	"Price": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Float),
	},
	"Image": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

//CreateProduct for creating a Product
var CreateProduct = &graphql.Field{
	Type:        ProductSchema, // the return type for this field
	Description: "Create new Product",
	Args:        ProductArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["Name"].(string)
		brand, okb := params.Args["Brand"]
		description, okd := params.Args["Description"]
		tags, okt := params.Args["Tags"]
		price, _ := params.Args["Price"].(float64)
		image, oki := params.Args["Image"]

		// perform mutation operation here
		// for e.g. create a Product and save to DB.
		newProduct := Product{
			Name:  name,
			Price: price,
		}
		if okb {
			newProduct.Brand = brand.(string)
		}
		if okd {
			newProduct.Description = description.(string)
		}
		if okt && tags != nil {
			DB.Where(tags.([]int)).Find(&newProduct.Tags)
		}
		if oki {
			newProduct.Image = image.(string)
		}
		DB.Create(&newProduct)

		// return the new Product object that we supposedly save to DB
		return newProduct, nil
	},
}

//UpdateProduct for updating a Product
var UpdateProduct = &graphql.Field{
	Type:        ProductSchema, // the return type for this field
	Description: "Update existing Product",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		name, okn := params.Args["Name"]
		brand, okb := params.Args["Brand"]
		description, okd := params.Args["Description"]
		addTag, okat := params.Args["AddTag"]
		removeTag, okrt := params.Args["RemoveTag"]
		price, okp := params.Args["Price"]
		image, oki := params.Args["Image"]

		// perform mutation operation here
		// will get the product from db
		// update the fields and update in the db
		product := Product{
			ID: uint(id),
		}
		tags := []Tag{}
		DB.Model(&product).Related(&tags, "Tags")
		product.Tags = tags

		if okn {
			product.Name = name.(string)
		}
		if okb {
			product.Brand = brand.(string)
		}
		if okd {
			product.Description = description.(string)
		}
		if okat {
			t := Tag{ID: uint(addTag.(int))}
			DB.First(&t)
			if len(t.Name) != 0 {
				product.Tags = append(product.Tags, t)
			}
		}
		if okrt {
			tag := Tag{ID: uint(removeTag.(int))}
			for i := len(product.Tags) - 1; i >= 0; i-- {
				if product.Tags[i].ID != tag.ID {
					continue
				}
				copy(product.Tags[i:], product.Tags[i+1:])
				product.Tags = product.Tags[:len(product.Tags)-1]
				DB.Model(&product).Association("Tags").Delete(&tag)
				break
			}
		}
		if okp {
			product.Price = price.(float64)
		}
		if oki {
			product.Image = image.(string)
		}
		DB.Save(&product)

		// return the new Product object that we supposedly save to DB
		return product, nil
	},
}

//ProductReview for adding areview to a Product
var ProductReview = &graphql.Field{
	Type:        ProductSchema, // the return type for this field
	Description: "Update existing Product",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"Review": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Rating": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"ReviewID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"Add": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"Remove": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		review, okr := params.Args["Review"]
		rating, okrr := params.Args["Rating"]
		reviewID, okri := params.Args["ReviewID"]
		add, oka := params.Args["Add"]
		remove, okre := params.Args["Remove"]

		// perform mutation operation here
		// will get the product from db
		// update the fields and update in the db
		product := Product{
			ID: uint(id),
		}
		reviews := []Review{}
		DB.Model(&product).Related(&reviews, "Reviews")
		product.Reviews = reviews

		if okr && okrr && oka {
			ad, ok := add.(bool)
			if ok && ad {
				r := Review{Review: review.(string), Rating: rating.(int)}
				DB.Create(&r)
				product.Reviews = append(product.Reviews, r)
			}
		}
		if okri && okre {
			rm, ok := remove.(bool)
			if rm && ok {
				review := Review{ID: uint(reviewID.(int))}
				for i := len(product.Reviews) - 1; i >= 0; i-- {
					if product.Reviews[i].ID != review.ID {
						continue
					}
					copy(product.Reviews[i:], product.Reviews[i+1:])
					product.Reviews = product.Reviews[:len(product.Reviews)-1]
					DB.Model(&product).Association("Reviews").Delete(&review)
					DB.Delete(&review)
					break
				}
			}
		}
		DB.Save(&product)

		// return the new Product object that we supposedly save to DB
		return product, nil
	},
}

//InitProduct will do the initializations required for the product
func InitProduct() {
	/*
	 * We will add ID, Add Tag, Remove Tag, Add Review and Remove Review to the update config
	 * and update the name and price config config
	 * delete the Tag
	 */
	args := graphql.FieldConfigArgument{}
	for k, v := range ProductArgumentConfig {
		args[k] = v
	}
	args["ID"] = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}
	args["AddTag"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["RemoveTag"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["AddReview"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["RemoveReview"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["Name"] = &graphql.ArgumentConfig{
		Type: graphql.String,
	}
	args["Price"] = &graphql.ArgumentConfig{
		Type: graphql.Float,
	}
	delete(args, "Tags")
	UpdateProduct.Args = args
}
