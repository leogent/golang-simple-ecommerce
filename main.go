package main

import (
	"fmt"
	"log"
	"simple-ecommerce/route"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Main Application Starts")

	//loading environmental variable
	loadenv()

	log.Fatal(route.RunApi("localhost:8090"))

}

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
