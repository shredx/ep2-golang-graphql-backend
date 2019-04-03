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
	Fields: &graphql.Fields{
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
