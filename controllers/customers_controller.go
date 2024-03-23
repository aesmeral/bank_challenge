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
		c.JSON(http.StatusNotFound, gin.H{"message": "Customer Not Found"})
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

func CustomerTransaction(c *gin.Context) {
	id := c.Params.ByName("id")

	var body struct {
		Value       int64
		Type        string
		Description string
	}
	var customer models.Customer

	c.Bind(&body)
	results := initializers.DB.First(&customer, id)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Customer Not Found"})
		return
	}

	if body.Value <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Value must be greater than 0"})
		return
	}

	if body.Type != "c" && body.Type != "d" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Only types c and d are accepted"})
		return
	}

	if len(body.Description) < 1 || len(body.Description) > 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Description must be between 1 and 10 characters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created"})
}
