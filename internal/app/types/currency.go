package types

import "time"

type Currency struct {
	From 		string		`db:"currency_from"`
	To 			string 		`db:"currency_to"`
	Well 		float64  	`db:"well"`
	UpdatedAt 	time.Time 	`db:"updated_at"`
}
