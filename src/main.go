package main

import (
	"github.com/jrschmidtt/api-gin/database"
	"github.com/jrschmidtt/api-gin/models"
	"github.com/jrschmidtt/api-gin/routes"
)

func main() {
	database.ConnectPostgresDB()
	models.Alunos = []models.Aluno{
		{Nome: "Aluno 1", CPF: "42942701013", RG: "47123453435"},
		{Nome: "Aluno 2", CPF: "98679221074", RG: "47123453435"},
		{Nome: "Aluno 3", CPF: "12863868020", RG: "47123453435"},
	}
	routes.HandleRequest()
}