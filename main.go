package main

import (
	"coinbase-go/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Todo struct to map the response
func main() {
	err := godotenv.Load()
	port := os.Getenv("PORT")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("succesfully load api")

	routes.InitRoutes()

	fmt.Println("server started at port:", port)
	http.ListenAndServe(":"+port, nil)
}
