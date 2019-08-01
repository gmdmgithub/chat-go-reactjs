package main

import (
	"chat-go-reactjs/pkg/websocket"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main router to serve pages
func router() http.Handler {

	pool := websocket.NewPool()
	go pool.Start()

	r := mux.NewRouter()
	// three way to implement handlefunc
	r.Path("/greeting").Methods(http.MethodGet).HandlerFunc(greet)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { //anonymous func
		log.Println("test")
		// http.head
		fmt.Fprintf(w, "just testing")
	})
	// three way to implement handlefunc

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	return r
}

// define our WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("new connection from host", r.Host)
	defer log.Println("websocket endpoint end", r.Host)

	// upgrade this connection to a WebSocket
	// connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Fatal("Upgrader problem", err)
	}

	client := &websocket.Client{
		ID:   clnID,
		Conn: conn,
		Pool: pool,
	}
	clnID = clnID + 1

	pool.Register <- client
	client.Read()
}

// homePage - just simple home page
func homePage(w http.ResponseWriter, req *http.Request) {
	state := "OK" //temporary everything is ok
	log.Print("request for status - let's see")
	fmt.Fprintf(w, state)
}

// greet - some response on greeting
func greet(w http.ResponseWriter, req *http.Request) {
	status, err := w.Write([]byte("Hello, world - nice to see you!"))
	if err != nil {
		log.Printf("problem with response %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Statuss is %v", status)
}
