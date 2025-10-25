package main

import (
	"fmt"
	"go-auth-manual/database"
	"go-auth-manual/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to homepage")
}

func main() {

	// env load
	err := godotenv.Load()
	if err != nil{
		log.Println(" .env not found")
	}

	// connect db first
	db := database.ConnectDB()
	defer db.Close()

	// creating handler, injecting dependecies
	h := handlers.NewHandler(db) // injecting dependecies

	// setup router
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/signup", h.SignUpHandler)
	mux.HandleFunc("/login", h.LoginHandler)

	fmt.Println("Server is starting on port :9000...")

	err = http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Println("Error starting server ", err)
	}
}
