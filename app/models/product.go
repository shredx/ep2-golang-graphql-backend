package models

import "github.com/jinzhu/gorm"

/*
 * This file contains the model defintions for Product
 */

//Product is the data structure of the product
type Product struct {
	gorm.Model
	Name        string   //Name of the product
	Brand       string   //Brand to which the product belongs to
	Description string   //Description of the product
	Tags        []string //Tags is the tags associated with the product
	Price       float32  //Price of the product
	Image       string   //Image url of the product
	Reviews     []Review //Reviews given to this product
}
