package models

import "time"

/*
 * This file contains the definitions of the order made by a user
 */

//Order made by the user
type Order struct {
	Cart           //Cart for which the order was made
	Date time.Time //Date on which the order was made
}
