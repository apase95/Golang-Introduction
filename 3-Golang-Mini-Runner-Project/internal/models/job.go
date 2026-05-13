package models

import "time"

type Job struct {
	ID 			uint 		`gorm:"primaryKey" json:"id"` 
	Name 		string 		`gorm:"size:50; not null" json:"name"`
	Command 	string		`gorm:"not null" json:"command"`
	Status 		string		`gorm:"default: 'PENDING'; not null" json:"status"`
	LogsPath 	string		`gorm:"" json:"logs_path,omitempty"`
	CreatedAt 	time.Time	`gorm:"" json:"created_at"`
	UpdatedAt 	time.Time	`gorm:"" json:"updated_at"`	
	FinishedAt 	*time.Time	`gorm:"" json:"finished_at,omitempty"`
}