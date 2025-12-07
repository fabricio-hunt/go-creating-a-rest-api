package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	Name  string
	Price float64
}

func main() {

	router := gin.Default()

	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Toscana, Calabresa, Margherita",
		})
	})

	router.Run() // listen and serve on :8080

	nomePizzaria := "Pizzaria GO"
	instagram, phone := "pizzariago", "1234-5678"

	println("Nome:", nomePizzaria)
	println("Informações da Pizzaria:")
	println("Instagram:", instagram)
	println("Telefone:", phone)
}
