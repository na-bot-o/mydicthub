package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/na-bot-o/rest/controllers"
)

func main() {

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/", controller.index()).Method("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
