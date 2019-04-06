package models

import (
	"github.com/graphql-go/graphql"
)

/*
 * This file conatins the defnitions of the shopping cart
 */

//Cart is the cart for shopping
type Cart struct {
	ID     uint    `gorm:"primary_key"`
	Items  []Item  `gorm:"foreign_key:CartID"` //Items is the list of items in the cart
	Total  float32 //Total price of items in the cart
	UserID uint
}

//CartConfig is the config for the Cart object
var CartConfig = graphql.ObjectConfig{
	Name: "Cart",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Items": &graphql.Field{
			Type: graphql.NewList(ItemSchema),
		},
		"Total": &graphql.Field{
			Type: graphql.Float,
		},
	},
}

//CartSchema is the schema for the Cart model
var CartSchema = graphql.NewObject(CartConfig)

//ReadCart will read a Cart given a CartID
var ReadCart = &graphql.Field{
	Type:        CartSchema, //the return of this field
	Description: "Get a single Cart and its detail",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		//marshall and cast the argument value
		id, ok := params.Args["ID"].(int)
		cart := Cart{ID: uint(id)}

		if ok {
			//Find the order from the DB
			DB.First(&cart).Related(&cart.Items, "Items")
		}
		//Getting the Products
		for i := 0; i < len(cart.Items); i++ {
			item := Item{ID: cart.Items[i].ID}
			//Getting the item from DB
			DB.First(&item).Related(&item.Product, "Product")
			cart.Items[i].Product = item.Product
		}
		// return the new order object that we supposedly have in DB
		return cart, nil
	},
}

//CreateCart for creating an Order
var CreateCart = &graphql.Field{
	Type:        CartSchema, // the return type for this field
	Description: "Create new Cart",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// perform mutation operation here
		// for e.g. create a Cart and save to DB.
		newCart := Cart{}

		//Creating the Cart in the DB
		DB.Create(&newCart)

		// return the new Product object that we supposedly save to DB
		return newCart, nil
	},
}

//UpdateCartArgumentConfig is argument config required update of the Order
var UpdateCartArgumentConfig = graphql.FieldConfigArgument{
	"ID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"Total": &graphql.ArgumentConfig{
		Type: graphql.Float,
	},
	"AddItem": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"RemoveItem": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

//UpdateCart is for updating an Order
var UpdateCart = &graphql.Field{
	Type:        CartSchema, //the return type for this field
	Description: "Update existing Cart",
	Args:        UpdateCartArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		//marshall and cast the argument
		id, _ := params.Args["ID"].(int)
		total, okt := params.Args["Total"]
		addItem, okai := params.Args["AddItem"]
		removeItem, okri := params.Args["RemoveItem"]

		// perform mutation operation here
		// will get the Cart from db
		// update the fields and update in the db
		cart := Cart{ID: uint(id)}

		if okt {
			cart.Total = total.(float32)
		}

		//Getting all the items
		items := []Item{}
		DB.Model(&cart).Related(&items, "Items")
		cart.Items = items

		if okai {
			i := Item{ID: uint(addItem.(int))}
			DB.First(&i)
			cart.Items = append(cart.Items, i)
		}

		if okri {
			item := Item{ID: uint(removeItem.(int))}
			for i := len(cart.Items) - 1; i >= 0; i++ {
				if cart.Items[i].ID != item.ID {
					continue
				}
				copy(cart.Items[i:], cart.Items[i+1:])
				cart.Items = cart.Items[:len(cart.Items)-1]
				DB.Model(&cart).Association("Items").Delete(&item)
				break
			}
		}

		DB.Save(&cart)

		//return the updated cart that we supposedly saved to the DB
		return cart, nil
	},
}
