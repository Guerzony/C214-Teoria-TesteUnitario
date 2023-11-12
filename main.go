package main

import (
	"C214-teoria-GO/database"
	"C214-teoria-GO/routes"
)

func main() {
	database.Conecta_BD()
	routes.HandleRequest()
}
