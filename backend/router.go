package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gmdmgithub/chat-go-reactjs/backend/pkg/websocket"

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
	conn, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader problem", err)
		return
	}
	client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
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

func writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			log.Println("Problem in nextreader message", err)
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println("Problem in nextwriter message", err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			log.Println("Problem in copy the writer to reader message", err)
			return
		}
		if err := w.Close(); err != nil {
			log.Println("Problem in closing the writer", err)
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
