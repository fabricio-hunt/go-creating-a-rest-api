package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Pizza represents a pizza item (model)
type Pizza struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {

	// Initialize Gin router
	router := gin.Default()

	// Register route to fetch pizzas
	router.GET("/pizzas", getPizzas)

	// Start server on default port :8080
	router.Run()

	// Basic pizzeria information (printed in terminal)
	name := "Pizzaria GO"
	instagram, phone := "pizzariago", "1234-5678"

	println("Name:", name)
	println("Pizzeria Info:")
	println("Instagram:", instagram)
	println("Phone:", phone)
}

// getPizzas handles GET requests and returns a list of pizzas in JSON format
func getPizzas(c *gin.Context) {

	// Example pizza list
	var pizzas = []Pizza{
		{ID: 1, Name: "Toscana", Price: 25.00},
		{ID: 2, Name: "Margherita", Price: 20.00},
		{ID: 3, Name: "Pepperoni", Price: 30.00},
	}
	c.JSON(http.StatusOK, gin.H{
		"pizzas": pizzas,
	})
}
