package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Value       int64    `json:"value"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	CustomerID  uint     `json:"-"`
	Customer    Customer `json:"customer,omitempty"`
}
