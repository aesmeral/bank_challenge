package controllers

import (
	"net/http"

	"github.com/aesmeral/bank_challenge/initializers"
	"github.com/aesmeral/bank_challenge/models"
	"github.com/gin-gonic/gin"
)

func CustomerIndex(c *gin.Context) {
	var customers []models.Customer
	initializers.DB.Find(&customers)

	results := make([]map[string]interface{}, len(customers))
	for index, customer := range customers {
		results[index] = customer.CustomerSerializer()
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func CustomerShow(c *gin.Context) {
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

func CustomerCreate(c *gin.Context) {
	var body struct {
		Balance int64
		Limit   int64
	}
	c.Bind(&body)

	customer := models.Customer{Limit: body.Limit, Balance: body.Balance}
	result := initializers.DB.Create(&customer)

	c.JSON(http.StatusOK, gin.H{
		"created": result.RowsAffected,
	})
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
