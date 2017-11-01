package model

import (

)

type Contact struct {
	First_name				string	`json:"first_name"`
	Last_name				string	`json:"last_name"`
	Email					string	`json:"email"`
	Street_address			string	`json:"street_address"`
	City					string	`json:"city"`
	State					string	`json:"state"`
	Zipcode					string	`json:"-"`
}
