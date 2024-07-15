package main

import "todo/frontend/server"

func main() {
	server.ImportFrontendHandlers()
	server.StartServer()
}
