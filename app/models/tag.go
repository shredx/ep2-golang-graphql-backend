package models

import "github.com/graphql-go/graphql"

/*
 * This file contains the model of tag
 */

//Tag is tag to for indexing purpose
type Tag struct {
	//ID of the Tag
	ID uint `gorm:"primary_key"`
	//Name of the tag
	Name string
}

//TagConfig is the object config for tag
var TagConfig = graphql.ObjectConfig{
	Name: "Tag",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
	},
}

//TagSchema is the schema of the Tag model
var TagSchema = graphql.NewObject(TagConfig)

//ReadTagResolve is the resolve function for both read tag/tags
var ReadTagResolve = func(params graphql.ResolveParams) (interface{}, error) {

	// marshall and cast the argument value
	id, ok := params.Args["ID"]
	tag := Tag{}

	if !ok {
		//id is not given
		var tags []Tag
		DB.Find(&tags)
		return tags, nil
	}

	//finding the Tag from the db
	DB.First(&tag, uint(id.(int)))

	// return the new Tag object that we supposedly save to DB
	return tag, nil
}

//ReadTag will read a Tag
var ReadTag = &graphql.Field{
	Type:        TagSchema, // the return type for this field
	Description: "Get a single Tag",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: ReadTagResolve,
}

//ReadTags will read a Tag
var ReadTags = &graphql.Field{
	Type:        graphql.NewList(TagSchema), // the return type for this field
	Description: "Get a single Tag",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: ReadTagResolve,
}

//CreateTag for creating a Tag
var CreateTag = &graphql.Field{
	Type:        TagSchema, // the return type for this field
	Description: "Create new Tag",
	Args: graphql.FieldConfigArgument{
		"Name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["Name"].(string)

		// perform mutation operation here
		// for e.g. create a Tag and save to DB.
		newTag := Tag{
			Name: name,
		}
		DB.Create(&newTag)

		// return the new Tag object that we supposedly save to DB
		return newTag, nil
	},
}

//UpdateTag for creating a Tag
var UpdateTag = &graphql.Field{
	Type:        TagSchema, // the return type for this field
	Description: "Update existing Tag",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"Name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["ID"].(int)
		name, _ := params.Args["Name"].(string)

		// perform mutation operation here
		// for e.g. create a Tag and save to DB.
		tag := Tag{
			ID:   uint(id),
			Name: name,
		}
		DB.Update(&tag)

		// return the new Tag object that we supposedly save to DB
		return tag, nil
	},
}
