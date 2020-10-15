package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taufikhidayat/pengenalan-microservice/menu-service/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-product", http.HandlerFunc(handler.AddMenuHandler))

	fmt.Println("Server listen on :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}
