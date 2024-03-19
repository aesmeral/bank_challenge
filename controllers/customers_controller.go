package controllers

import (
	"net/http"

	"github.com/aesmeral/bank_challenge/initializers"
	"github.com/aesmeral/bank_challenge/models"
	"github.com/gin-gonic/gin"
)

func CustomerIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "we're in the customers",
	})
}

func ShowCustomer(c *gin.Context) {
	id := c.Params.ByName("id")

	var customer models.Customer
	results := initializers.DB.First(&customer, id)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	data := gin.H{"data": customer.CustomerSerializer()}

	c.JSON(http.StatusOK, data)
}

func CustomerCreateInit(c *gin.Context) {
	customers := []models.Customer{
		{Limit: 100000, Balance: 0},
		{Limit: 80000, Balance: 0},
		{Limit: 1000000, Balance: 0},
		{Limit: 10000000, Balance: 0},
		{Limit: 500000, Balance: 0},
	}

	result := initializers.DB.Create(customers)

	c.JSON(http.StatusOK, gin.H{
		"created": result.RowsAffected,
	})
}
