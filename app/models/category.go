package models

import (
	"github.com/graphql-go/graphql"
)

/*
 * This file contains the defnitions of the category model
 */

//Category helps to classify the products.
//It can have subcategories too
type Category struct {
	//ID of the product
	ID          uint   `gorm:"primary_key"`
	Name        string //Name of the category
	Description string //Description is the description for the category
	//SubCategories for further classification
	SubCategories []Category `gorm:"many2many:category_subcategories;association_jointable_foreignkey:subcategory_id"`
	//Products has the list of products in a category
	Products []Product `gorm:"many2many:category_products;"`
}

//CategoryConfig is the config for the category object
var CategoryConfig = graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"Description": &graphql.Field{
			Type: graphql.String,
		},
		//Subcategories will be added in init categories
		"Products": &graphql.Field{
			Type: graphql.NewList(ProductSchema),
		},
	},
}

//CategorySchema is the schema of the category model
var CategorySchema = graphql.NewObject(CategoryConfig)

//ReadCategoryResolve is the resolve function for the category read resolve
var ReadCategoryResolve = func(params graphql.ResolveParams) (interface{}, error) {

	// marshall and cast the argument value
	id, ok := params.Args["ID"].(int)
	category := Category{ID: uint(id)}

	if !ok {
		cats := []Category{}
		DB.Find(&cats)
		for i := 0; i < len(cats); i++ {
			DB.First(&cats[i]).Related(&cats[i].Products, "Products").Related(&cats[i].SubCategories, "SubCategories")
		}
		return cats, nil
	}

	//finding the Product from the db
	DB.First(&category).Related(&category.Products, "Products").Related(&category.SubCategories, "SubCategories")

	// return the new Product object that we supposedly save to DB
	return category, nil
}

//ReadCategory will read a Category
var ReadCategory = &graphql.Field{
	Type:        CategorySchema, // the return type for this field
	Description: "Get a single category",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: ReadCategoryResolve,
}

//ReadCategories will read all categories
var ReadCategories = &graphql.Field{
	Type:        graphql.NewList(CategorySchema), // the return type for this field
	Description: "Get all the categories",
	Resolve:     ReadCategoryResolve,
}

//CategoryArgumentConfig is argument config required for the category
var CategoryArgumentConfig = graphql.FieldConfigArgument{
	"Name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"Description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

//CreateCategory for creating a Category
var CreateCategory = &graphql.Field{
	Type:        CategorySchema, // the return type for this field
	Description: "Create new Category",
	Args:        CategoryArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["Name"].(string)
		description, okd := params.Args["Description"]

		// perform mutation operation here
		// for e.g. create a Category and save to DB.
		newCategory := Category{
			Name: name,
		}
		if okd {
			newCategory.Description = description.(string)
		}
		DB.Create(&newCategory)

		// return the new Product object that we supposedly save to DB
		return newCategory, nil
	},
}

//UpdateCategory for creating a Category
var UpdateCategory = &graphql.Field{
	Type:        CategorySchema, // the return type for this field
	Description: "Update existing Category",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		name, okn := params.Args["Name"]
		description, okd := params.Args["Description"]
		addCategory, okas := params.Args["AddSubCategory"]
		removeCategory, okrs := params.Args["RemoveSubCategory"]
		addProduct, okap := params.Args["AddProduct"]
		removeProduct, okrp := params.Args["RemoveProduct"]

		// perform mutation operation here
		// will get the product from db
		// update the fields and update in the db
		category := Category{
			ID: uint(id),
		}
		DB.First(&category)
		DB.Model(&category).Related(&category.SubCategories, "SubCategories").Related(&category.Products, "Products")

		if okn {
			category.Name = name.(string)
		}
		if okd {
			category.Description = description.(string)
		}
		if okas {
			sc := Category{ID: uint(addCategory.(int))}
			DB.First(&sc)
			if len(sc.Name) != 0 {
				category.SubCategories = append(category.SubCategories, sc)
			}
		}
		if okrs {
			sc := Category{ID: uint(removeCategory.(int))}
			for i := len(category.SubCategories) - 1; i >= 0; i-- {
				if category.SubCategories[i].ID != sc.ID {
					continue
				}
				copy(category.SubCategories[i:], category.SubCategories[i+1:])
				category.SubCategories = category.SubCategories[:len(category.SubCategories)-1]
				DB.Model(&category).Association("SubCategories").Delete(&sc)
				break
			}
		}
		if okap {
			p := Product{ID: uint(addProduct.(int))}
			DB.First(&p)
			if len(p.Name) != 0 {
				category.Products = append(category.Products, p)
			}
		}
		if okrp {
			p := Product{ID: uint(removeProduct.(int))}
			for i := len(category.Products) - 1; i >= 0; i-- {
				if category.Products[i].ID != p.ID {
					continue
				}
				copy(category.Products[i:], category.Products[i+1:])
				category.Products = category.Products[:len(category.Products)-1]
				DB.Model(&category).Association("Products").Delete(&p)
				break
			}
		}
		DB.Save(&category)

		// return the new Product object that we supposedly save to DB
		return category, nil
	},
}

//InitCategories will do the required initalizations for the category model to work with graphql
func InitCategories() {
	/*
	 * We will add the Subcategories to the category config
	 * Then we will init the args fopr the update category
	 */
	//Adding the subscategories to the categories schema
	CategorySchema.AddFieldConfig("SubCategories", &graphql.Field{
		Type: graphql.NewList(CategorySchema),
	})

	args := graphql.FieldConfigArgument{}
	for k, v := range CategoryArgumentConfig {
		args[k] = v
	}
	args["ID"] = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}
	args["AddSubCategory"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["RemoveSubCategory"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["AddProduct"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["RemoveProduct"] = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}
	args["Name"] = &graphql.ArgumentConfig{
		Type: graphql.String,
	}
	UpdateCategory.Args = args
}
