package main

import (
	"github.com/aesmeral/bank_challenge/initializers"
	"github.com/aesmeral/bank_challenge/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.CustomerRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
