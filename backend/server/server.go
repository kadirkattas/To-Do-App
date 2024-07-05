package server

import (
	"log"
	"net/http"

	"todo/backend/auth"
	logincontroller "todo/backend/controllers/login"
	registercontroller "todo/backend/controllers/register"

	"github.com/gorilla/mux"
)

const backEndPort = ":8080"

var router = mux.NewRouter()

func StartServer() {
	log.Println("Server is running on port " + backEndPort)
	err := http.ListenAndServe(backEndPort, router)
	if err != nil {
		log.Fatal(err)
	}
}

func ImportHandlers() {
	router.HandleFunc("/api/v1/user/register", registercontroller.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/user/login", logincontroller.LoginHandler).Methods("POST")

	protectedRoutes := router.PathPrefix("/api/v1/user").Subrouter()
	protectedRoutes.Use(auth.AuthMiddleware)
}
