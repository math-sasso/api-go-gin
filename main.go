package main

import (
	"github.com/math-sasso/api-go-gin/database"
	"github.com/math-sasso/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
