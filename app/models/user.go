package models

import (
	"github.com/graphql-go/graphql"
)

/*
 * This file contains the defintions of user model
 */

//User is the model of the user
type User struct {
	ID       uint   `gorm:"primary_key"` //ID of the user
	Number   string //Number is the mobile number of the user
	Password string //Password is the password the user
	Name     string //Name is the display name of the user
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
	},
}

//UserSchema is the schema of the user model
var UserSchema = graphql.NewObject(UserConfig)

//ReadUserResolve is the resolve function for user and list of users
var ReadUserResolve = func(params graphql.ResolveParams) (interface{}, error) {

	// marshall and cast the argument value
	id, ok := params.Args["ID"].(int)
	user := User{ID: uint(id)}

	if !ok {
		//return all the users
		var users []User
		DB.Find(&users)
		return users, nil
	}

	//finding the user from the db
	DB.First(&user)

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
	Description: "Get a single user",
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

//UpdateUser for creating a user
var UpdateUser = &graphql.Field{
	Type:        UserSchema, // the return type for this field
	Description: "Update existing user",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
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
		id, _ := params.Args["ID"].(int)
		number, _ := params.Args["Number"].(string)
		password, _ := params.Args["Password"].(string)
		name, _ := params.Args["Name"].(string)

		// perform mutation operation here
		// for e.g. update the user and save to DB.
		user := User{
			ID:       uint(id),
			Number:   number,
			Password: password,
			Name:     name,
		}
		DB.Update(&user)

		//Emptying the password
		user.Password = ""

		// return the new User object that we supposedly save to DB
		return user, nil
	},
}
