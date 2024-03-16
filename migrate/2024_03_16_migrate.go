package main

import (
	"github.com/aesmeral/bank_challenge/initializers"
	"github.com/aesmeral/bank_challenge/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	db := initializers.DB

	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Transaction{})
}
