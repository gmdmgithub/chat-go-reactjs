package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port = "8080"

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
