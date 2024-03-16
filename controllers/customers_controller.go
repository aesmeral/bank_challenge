package controllers

import "github.com/gin-gonic/gin"

func CustomerIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "we're in the customers",
	})
}
