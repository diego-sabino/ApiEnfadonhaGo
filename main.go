package main

import (
	"mymodule/database"
	"mymodule/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
