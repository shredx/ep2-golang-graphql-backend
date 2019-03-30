package models

import "github.com/jinzhu/gorm"

/*
 * This file conatins the defnitions of the shopping cart
 */

//Cart is the cart for shopping
type Cart struct {
	gorm.Model
	Items []Item  //Items is the list of items in the cart
	Total float32 //Total price of items in the cart
}
