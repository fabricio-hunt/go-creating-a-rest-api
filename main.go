package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// pizzas stores an in-memory list of pizza items.
// In real applications, this would be replaced by a database.
var pizzas []models.Pizza

func main() {

	// Pre-load some pizza data
	loadPizzas()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasByID)

	// Start server on default port :8080
	router.Run()

	// Basic pizzeria information (printed to console)
	name := "Pizzaria GO"
	instagram := "pizzariago"
	phone := "1234-5678"

	fmt.Println("Name:", name)
	fmt.Println("Pizzeria Info:")
	fmt.Println("Instagram:", instagram)
	fmt.Println("Phone:", phone)
}

// getPizzas handles GET requests and returns a list of pizzas in JSON format.
func getPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": pizzas,
	})
}

// postPizzas handles POST requests to add a new pizza.
// It binds the received JSON to a Pizza struct.
func postPizzas(c *gin.Context) {
	var newPizza models.Pizza

	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newPizza.ID = len(pizzas) + 1 // Simple ID assignment
	savePizzas()                  // Save updated pizzas to file

	// Add new pizza to the in-memory list
	pizzas = append(pizzas, newPizza)
	c.JSON(http.StatusCreated, newPizza)
}

// getPizzasByID handles GET requests to retrieve a pizza by its ID.
func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "pizza not found"})
}

func loadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error opening pizzas.json:", err)
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding pizzas.json:", err)
	}
}

func savePizzas() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error creating pizzas.json:", err)
		return
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding pizzas.json:", err)
	}
}
