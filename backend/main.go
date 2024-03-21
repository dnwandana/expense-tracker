package main

import (
	"log"
	"net/http"

	"github.com/dnwandana/expense-tracker/database"
	"github.com/dnwandana/expense-tracker/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("setting up database connection")
	db := database.NewConnection()

	log.Println("setting up routes")
	mux := http.NewServeMux()
	routes.SetupAuthenticationRoutes(mux, db)
	routes.SetupCategoryRoutes(mux, db)

	server := http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	log.Println("server up and running on port: 5000")
	log.Fatal(server.ListenAndServe())
}
