package main

import (
	"todo/backend/database"
	"todo/backend/server"
)

func main() {
	server.ImportHandlers()

	database.InitializeDatabase()

	server.StartServer()
}
