package server

import (
	"log"
	"net/http"

	loginpage "todo/frontend/static/loginPage"
)

const frontEndPort = ":8081"

func StartServer() {
	log.Println("Server is running on port " + frontEndPort)
	err := http.ListenAndServe(frontEndPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ImportFrontendHandlers() {
	http.HandleFunc("/login", loginpage.LoginPageHandler)
}
