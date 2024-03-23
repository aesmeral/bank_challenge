package routes

import (
	"github.com/aesmeral/bank_challenge/controllers"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	customers := r.Group("/customers")
	{
		customers.GET("/", controllers.CustomerIndex)
		customers.GET("/:id", controllers.CustomerShow)
		customers.POST("/", controllers.CustomerCreate)
		customers.POST("/:id/transaction", controllers.CustomerTransaction)
	}
}
