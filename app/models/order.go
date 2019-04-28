package models

import (
	"time"

	"github.com/graphql-go/graphql"
)

/*
 * This file contains the definitions of the order made by a user
 */

//Order made by the user
type Order struct {
	ID     uint      `gorm:"primary_key"`
	Name   string    //The Order name to be displayed in the frontend
	Items  []Item    `gorm:"foreign_key:OrderID"` //Items is the list of items in the cart
	Total  float64   //Total price of items in the cart
	Date   time.Time //Date on which the order was made
	UserID uint
}

//OrderConfig is the config for the Order object
var OrderConfig = graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"Items": &graphql.Field{
			Type:        graphql.NewList(ItemSchema),
			Description: "Getting the list of items",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				//Getting the source Order
				order, ok := params.Source.(Order)
				if !ok {
					return nil, nil
				}
				//Getting the Items from DB
				DB.First(&order).Related(&order.Items, "Items")
				return order.Items, nil
			},
		},
		"Total": &graphql.Field{
			Type: graphql.Float,
		},
		"Date": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
}

//OrderSchema is the schema of the Order model
var OrderSchema = graphql.NewObject(OrderConfig)

//ReadOrder will read a Order given an OrderID
var ReadOrder = &graphql.Field{
	Type:        OrderSchema, //the return of this field
	Description: "Get a single Order and its detail",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		//marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		order := Order{ID: uint(id)}
		//Finding the order from DB
		DB.First(&order)
		// return the new order object that we supposedly have in DB
		return order, nil
	},
}

//CreateOrderArgumentConfig is argument config required for the Order
var CreateOrderArgumentConfig = graphql.FieldConfigArgument{
	"Name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"Total": &graphql.ArgumentConfig{
		Type: graphql.Float,
	},
	"Date": &graphql.ArgumentConfig{
		Type: graphql.DateTime,
	},
}

//CreateOrder for creating an Order
var CreateOrder = &graphql.Field{
	Type:        OrderSchema, // the return type for this field
	Description: "Create new Order",
	Args:        CreateOrderArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["Name"].(string)
		total, okt := params.Args["Total"]
		date, okd := params.Args["Date"]

		// perform mutation operation here
		// for e.g. create an Order and save to DB.
		newOrder := Order{Name: name}

		if okt {
			newOrder.Total = total.(float64)
		}
		if okd {
			newOrder.Date = date.(time.Time)
		} else {
			//If Date not given then adding today's date
			newOrder.Date = time.Now()
		}

		//Creating the Order in the DB
		DB.Create(&newOrder)

		// return the new Product object that we supposedly save to DB
		return newOrder, nil
	},
}

//UpdateOrderArgumentConfig is argument config required update of the Order
var UpdateOrderArgumentConfig = graphql.FieldConfigArgument{
	"ID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"Name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"Total": &graphql.ArgumentConfig{
		Type: graphql.Float,
	},
	"Date": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.DateTime), //The format of DateTime at frontend is "2017-10-06T03:40:00.000Z"
	},
	"AddItem": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"RemoveItem": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

//UpdateOrder is for updating an Order
var UpdateOrder = &graphql.Field{
	Type:        OrderSchema, //the return type for this field
	Description: "Update existing Order",
	Args:        UpdateOrderArgumentConfig,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		//marshall and cast the argument
		id, _ := params.Args["ID"].(int)
		name, okn := params.Args["Name"]
		total, okt := params.Args["Total"]
		date, _ := params.Args["Date"]
		addItem, okai := params.Args["AddItem"]
		removeItem, okri := params.Args["RemoveItem"]

		// perform mutation operation here
		// will get the Order from db
		// update the fields and update in the db
		order := Order{
			ID:   uint(id),
			Date: date.(time.Time),
		}

		if okn {
			order.Name = name.(string)
		}
		if okt {
			order.Total = total.(float64)
		}

		//Getting all the items
		items := []Item{}
		DB.Model(&order).Related(&items, "Items")
		order.Items = items

		if okai {
			i := Item{ID: uint(addItem.(int))}
			DB.First(&i)
			order.Items = append(order.Items, i)
		}

		if okri {
			item := Item{ID: uint(removeItem.(int))}
			for i := len(order.Items) - 1; i >= 0; i++ {
				if order.Items[i].ID != item.ID {
					continue
				}
				copy(order.Items[i:], order.Items[i+1:])
				order.Items = order.Items[:len(order.Items)-1]
				DB.Model(&order).Association("Items").Delete(&item)
				break
			}
		}

		DB.Save(&order)

		//return the updated order that we supposedly saved to the DB
		return order, nil
	},
}
