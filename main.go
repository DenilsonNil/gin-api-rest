package main

import (
	"github.com/DenilsonNil/gin-api-rest/database"
	"github.com/DenilsonNil/gin-api-rest/routes"
)

func main() {
	database.DbConnect()
	// models.Alunos = []models.Aluno{
	// 	{ID: "1", Name: "Denilson", RG: "123456789", CPF: "98765432100"},
	// 	{ID: "2", Name: "Maria", RG: "987654321", CPF: "12345678900"},
	// 	{ID: "3", Name: "João", RG: "456789123", CPF: "32165498700"},
	// }
	routes.HandleRequests()
}
