package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hafrizresa/contacts/controllers"
)

func main() {
	router := mux.NewRouter()

	//INI MIDDLEWARE
	// router.Use()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println(port)

	router.HandleFunc("/api/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/login", controllers.AuthUser).Methods("POST")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
