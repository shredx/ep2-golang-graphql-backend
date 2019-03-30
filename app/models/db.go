package models

import "github.com/jinzhu/gorm"

/*
 This file contains the variables and utilities required for the DB interaction
*/

//DB is the data base connection instance to be accessed globally
var DB *gorm.DB
