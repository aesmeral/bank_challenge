package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID      uint  `json:"id"`
	Limit   int64 `json:"limit"`
	Balance int64 `json:"balance"`
}
