package main

import (
	"github.com/gorilla/mux"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/user", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/user/login", handler.Auth).Methods(http.MethodPost)
	router.Use(middleware.Logging)
	router.Use(middleware.Recover)
	router.Use(middleware.Accept)

	log.Fatal(http.ListenAndServe(":8080", router))
}
