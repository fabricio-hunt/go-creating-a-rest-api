package main

//principal function

type Pizza struct {
	ID    int
	name  string
	price float64
}

func main() {

	//variable declaration style short	assignment
	nomePizzaria := "Pizzaria GO"
	instagram, phone := "pizzariago", "1234-5678"

	//print values
	println("Nome:", nomePizzaria)
	println("Informações da Pizzaria:")
	println("Instagram:", instagram)
	println("Telefone:", phone)
}
