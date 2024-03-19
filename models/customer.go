package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Limit   int64 `json:"limit"`
	Balance int64 `json:"balance"`
}

func (c Customer) CustomerSerializer() map[string]interface{} {
	serializer := map[string]interface{}{
		"id":      c.ID,
		"limit":   c.Limit,
		"balance": c.Balance,
	}

	return serializer
}
