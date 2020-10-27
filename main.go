package main

import (
	"coinbase-go/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// Todo struct to map the response
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("succesfully load api")

	routes.InitRoutes()
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
