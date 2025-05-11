package main

import (
	"GoProjects/database"
	"GoProjects/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
