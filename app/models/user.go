package models

import "github.com/jinzhu/gorm"

/*
 * This file contains the defintions of user model
 */

//User is the model of the user
type User struct {
	gorm.Model
	Number   string //Number is the mobile number of the user
	Password string //Password is the password the user
	Name     string //Name is the display name of the user
}
