package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
)

// It is need to define an Upgrader, this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// main router to serve pages
func router() http.Handler {
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

	r.HandleFunc("/ws", serveWs)

	return r
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection from host", r.Host)
	defer log.Println("websocket endpoint end", r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader problem", err)
		return
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

// define a reader which will listen for
// new messages being sent to our WebSocket endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Problem in reading message", err)
			return
		}
		// print out that message for clarity
		log.Printf("Message is processed %s and type is %v", string(p), messageType)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Problem in writing message", err)
			return
		}

	}
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
