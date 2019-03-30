package models

import "github.com/jinzhu/gorm"

/*
 * This file contains the defnitions of the category model
 */

//Category helps to classify the products.
//It can have subcategories too
type Category struct {
	gorm.Model
	Name          string     //Name of the category
	Description   string     //Description is the description for the category
	SubCategories []Category //SubCategories for further classification
	Products      []Product  //Products has the list of products in a category
}
