package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

)

var port string

func init(){

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
	log.Printf("Main started")
	defer log.Printf("End of main!")

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router()))

}

func router() http.Handler {
	r := mux.NewRouter()
	r.Path("/greeting").Methods(http.MethodGet).HandlerFunc(greet)

	r.HandleFunc("/", homePage)

	return r
}

func homePage(w http.ResponseWriter, req *http.Request) {
	state := "OK" //temporary everything is ok
	log.Print("request for status - let's see")
	fmt.Fprintf(w, state)
}
func greet(w http.ResponseWriter, req *http.Request) {
	status, err := w.Write([]byte("Hello, world - nice to see you!"))
	if err != nil {
		log.Printf("problem with response %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Statuss is %v", status)
}
