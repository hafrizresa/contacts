package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hafrizresa/contacts/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	router.HandleFunc(controllers.CreateAccount).Methods(("GET"))
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
