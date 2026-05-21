package models

import "time"

type Expense struct {
	ID 			int 		`json:"id"`
	Amount 		float64 	`json:"amount"`
	Note 		string 		`json:"note"`
	CreatedAd 	time.Time 	`json:"created_at"`
}