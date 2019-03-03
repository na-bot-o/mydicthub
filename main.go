package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/na-bot-o/mydicthub/controllers"
)

func main() {

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/", controller.IndexHandler).Methods("GET")
	router.HandleFunc("/login", controller.LoginHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8085", router))
}
