package models

import "github.com/jinzhu/gorm"

/*
 * This file contains the model of review
 */

//Review given by a user for the product
type Review struct {
	gorm.Model
	//Review of the user
	Review string
	//Rating given by the user
	Rating int
}
