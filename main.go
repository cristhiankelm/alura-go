package main

import (
	"GoProjects/database"
	"GoProjects/models"
	"GoProjects/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Lucas", CPF: "123456789", RG: "987654321"},
		{Nome: "Ana", CPF: "987654321", RG: "123456789"},
	}
	routes.HandleRequests()
}
