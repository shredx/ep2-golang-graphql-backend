package models

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

/*
 * This file contains the defintions of user model
 */

//User is the model of the user
type User struct {
	ID       uint    `gorm:"primary_key"` //ID of the user
	Number   string  //Number is the mobile number of the user
	Password string  //Password is the password the user
	Name     string  //Name is the display name of the user
	Cart     Cart    `gorm:"foreign_key:UserID"` //This is the current cart of the user
	Orders   []Order `gorm:"foreign_key:UserID"` //Orders given by this user
}

//UserConfig is the object config for user
var UserConfig = graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Number": &graphql.Field{
			Type: graphql.String,
		},
		"Password": &graphql.Field{
			Type: graphql.String,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"Cart": &graphql.Field{
			Type: CartSchema,
		},
		"Orders": &graphql.Field{
			Type: graphql.NewList(OrderSchema),
		},
	},
}

//UserSchema is the schema of the user model
var UserSchema = graphql.NewObject(UserConfig)

//ReadUserResolve is the resolve function for user and list of users
var ReadUserResolve = func(params graphql.ResolveParams) (interface{}, error) {

	// marshall and cast the argument value
	id, ok := params.Args["ID"].(int)
	user := User{ID: uint(id)}

	//Getting the Items in Cart
	DB.First(&user.Cart).Related(&user.Cart.Items, "Items")

	if !ok {
		//return all the users
		users := []User{}
		DB.Find(&users)
		for i := 0; i < len(users); i++ {
			DB.Find(&users[i]).Related(&user.Cart, "Cart").Related(&user.Orders, "Orders")
			//Emptying the password
			users[i].Password = ""

			//Getting the Items in Cart
			cart := Cart{ID: users[i].Cart.ID}
			DB.First(&cart).Related(&cart.Items, "Items")
			fmt.Println(cart)
			users[i].Cart.Items = cart.Items
		}

		return users, nil
	}

	//finding the user from the db
	DB.First(&user).Related(&user.Cart, "Cart").Related(&user.Orders, "Orders")

	//Emptying the password
	user.Password = ""

	// return the new User object that we supposedly save to DB
	return user, nil
}

//ReadUser will read a user
var ReadUser = &graphql.Field{
	Type:        UserSchema, // the return type for this field
	Description: "Get a single user",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: ReadUserResolve,
}

//ReadUsers will read a user
var ReadUsers = &graphql.Field{
	Type:        graphql.NewList(UserSchema), // the return type for this field
	Description: "Get all users",
	Resolve:     ReadUserResolve,
}

//CreateUser for creating a user
var CreateUser = &graphql.Field{
	Type:        UserSchema, // the return type for this field
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"Number": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		number, _ := params.Args["Number"].(string)
		password, _ := params.Args["Password"].(string)
		name, _ := params.Args["Name"].(string)

		// perform mutation operation here
		// for e.g. create a User and save to DB.
		newUser := User{
			Number:   number,
			Password: password,
			Name:     name,
		}
		DB.Create(&newUser)

		//Emptying the password
		newUser.Password = ""

		// return the new User object that we supposedly save to DB
		return newUser, nil
	},
}

//UpdateUser for updating a user
var UpdateUser = &graphql.Field{
	Type:        UserSchema, // the return type for this field
	Description: "Update existing user",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"Number": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"AddOrder": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"RemoveOrder": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"AddCart": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"RemoveCart": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		number, okn := params.Args["Number"]
		password, _ := params.Args["Password"].(string)
		name, _ := params.Args["Name"].(string)
		addOrder, okao := params.Args["AddOrder"]
		removeOrder, okro := params.Args["RemoveOrder"]
		addCart, okac := params.Args["AddCart"]
		removeCart, okrc := params.Args["RemoveCart"]

		// perform mutation operation here
		// for e.g. update the user and save to DB.
		// will get the User from db
		// update the fields and update in the db
		user := User{
			ID:       uint(id),
			Password: password,
			Name:     name,
		}

		if okn {
			user.Number = number.(string)
		}

		//Getting the list of Orders
		orders := []Order{}
		DB.Model(&user).Related(&orders, "Orders")
		user.Orders = orders

		if okao {
			o := Order{ID: uint(addOrder.(int))}
			DB.First(&o)
			user.Orders = append(user.Orders, o)
		}

		if okro {
			o := Order{ID: uint(removeOrder.(int))}
			for i := len(user.Orders) - 1; i >= 0; i++ {
				if user.Orders[i].ID != o.ID {
					continue
				}
				copy(user.Orders[i:], user.Orders[i+1:])
				user.Orders = user.Orders[:len(user.Orders)-1]
				DB.Model(&user).Association("Orders").Delete(&o)
				break
			}
		}

		if okac {
			c := Cart{ID: uint(addCart.(int))}
			DB.First(&c)
			user.Cart = c
		}

		if okrc {
			cart := Cart{}
			DB.Where(removeCart.(int)).Find(&cart)
			DB.Model(&user).Association("Cart").Delete(&user.Cart)
			user.Cart = Cart{}
		}

		DB.Save(&user)

		//Emptying the password
		user.Password = ""

		// return the new User object that we supposedly save to DB
		return user, nil
	},
}
