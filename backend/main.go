package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var port string

var clnID int

func init() {
	log.Println("Init main")

	clnID = 1

	// first read .env file and put it to env
	if err := godotenv.Load(); err != nil {
		log.Printf("Fatal problem during initialization: %v\n", err)
		os.Exit(1)
	}

	p, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		log.Print("No http port in .env file, default 8000 taken")
		p = ":8080"
	}

	port = p

}

func main() {
	log.Println("Main started")
	defer log.Printf("End of main!")

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router()))

}
