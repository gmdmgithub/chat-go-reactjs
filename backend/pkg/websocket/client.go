package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//Client struct to keep client data
type Client struct {
	ID   string          `json:"id,omitempty"`
	Conn *websocket.Conn `json:"conn,omitempty"`
	Pool *Pool           `json:"pool,omitempty"`
}

// Message - struct to keep messages
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

//Read - method to read client messages
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
