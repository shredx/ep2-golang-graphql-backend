package models

import (
	"github.com/graphql-go/graphql"
)

/*
 * This file contains the defintions for the item which is a basic unit in the cart with qty
 */

//Item is the basic unit in the cart the identifies a product with price and qty
type Item struct {
	ID      uint    `gorm:"primary_key"`
	Product Product `gorm:"foreign_key:ItemID"` //Product the item represents
	Price   float32 //Price is the price of a single item
	Qty     int     //Qty is the number of items purchased
	OrderID uint    //This is the foreign_key for Order.Items
	CartID  uint    //This is the foreign_key for Cart.Items
}

//ItemConfig is the config for the Item object
var ItemConfig = graphql.ObjectConfig{
	Name: "Item",
	Fields: &graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Product": &graphql.Field{
			Type: graphql.NewNonNull(ProductSchema),
		},
		"Price": &graphql.Field{
			Type: graphql.Float,
		},
		"Qty": &graphql.Field{
			Type: graphql.Float,
		},
	},
}

//ItemSchema is the schema for the Item model
var ItemSchema = graphql.NewObject(ItemConfig)

//ReadItem will read a Item given an ItemID
var ReadItem = &graphql.Field{
	Type:        ItemSchema, //the return of this field
	Description: "Get a single Item and its detail",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		//marshall and cast the argument value
		id, ok := params.Args["ID"].(int)
		item := Item{ID: uint(id)}

		if ok {
			//Find the order from the DB
			DB.First(&item).Related(&item.Product, "Product")
		}
		// return the new item object that we supposedly have in DB
		return item, nil
	},
}

//ItemArgumentConfig is argument config required for the Item
var ItemArgumentConfig = graphql.FieldConfigArgument{
	"Product": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"Price": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Float),
	},
	"Qty": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

//CreateItem for creating an Item
var CreateItem = &graphql.Field{
	Type:        ItemSchema, // the return type for this field
	Description: "Create new Item",
	Args:        ItemArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		p, _ := params.Args["Product"]
		price, _ := params.Args["Price"]
		qty, _ := params.Args["Qty"]

		// perform mutation operation here
		// for e.g. create an Item and save to DB.
		newItem := Item{
			Price: price.(float32),
			Qty:   qty.(int),
		}

		//Getting the Product from the DB
		DB.Where(p.(int)).Find(&newItem.Product)

		//Creating the Item in the DB
		DB.Create(&newItem)

		// return the new Item object that we supposedly save to DB
		return newItem, nil
	},
}

//UpdateItemArgumentConfig is argument config required for the Item
var UpdateItemArgumentConfig = graphql.FieldConfigArgument{
	"ID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"Price": &graphql.ArgumentConfig{
		Type: graphql.Float,
	},
	"Qty": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"ChangeProduct": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

//UpdateItem for updating an Item
var UpdateItem = &graphql.Field{
	Type:        ItemSchema, // the return type for this field
	Description: "Update an Item",
	Args:        UpdateItemArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		price, okp := params.Args["Price"]
		qty, okq := params.Args["Qty"]
		changeProduct, okcp := params.Args["ChangeProduct"]

		// perform mutation operation here
		// for e.g. update an Item and save to DB.
		// will get the Item from db
		// update the fields and update in the db
		item := Item{ID: uint(id)}

		if okp {
			item.Price = price.(float32)
		}
		if okq {
			item.Qty = qty.(int)
		}

		//Getting the item from DB
		DB.First(&item).Related(&item.Product, "Product")

		if okcp {
			//Getting the Product from the DB
			pdt := Product{}
			DB.Where(changeProduct.(int)).Find(&pdt)
			DB.Model(&item).Association("Product").Delete(&item.Product)
			item.Product = pdt
		}

		//Updating the Item in the DB
		DB.Save(&item)

		// return the new Item object that we supposedly save to DB
		return item, nil
	},
}
