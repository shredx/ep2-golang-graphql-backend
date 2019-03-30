package models

import "github.com/jinzhu/gorm"

/*
 * This file contains the defintions for the item which is a basic unit in the cart with qty
 */

//Item is the basic unit in the cart the identifies a product with price and qty
type Item struct {
	gorm.Model
	Product Product //Product the item represents
	Price   float32 //Price is the price of a single item
	Qty     int     //Qty is the number of items purchased
}
