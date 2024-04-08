package main

import (
	"Teste/Alura/ApiRest/database"
	"Teste/Alura/ApiRest/models"
	"Teste/Alura/ApiRest/routes"
	"fmt"
)

func main() {

	// exemplo de TESTE
	// mocado (quando não está no Banco de dados)
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Claudio", Historia: "Historia 1"},
		{Id: 2, Nome: "Renata", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor com GO")
	routes.HandleRequest()
}
