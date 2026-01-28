package main

import (
	"log"
	"net/http"

	"crud/db"
	"crud/handlers"
)

func main() {
	db.Connect()

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UserHandler)

	log.Println("Server running on :8080")
	log.Error(http.ListenAndServe(":8080", nil))
}
