package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//Client struct to keep client data
type Client struct {
	ID   int             `json:"id,omitempty"`
	Conn *websocket.Conn `json:"conn,omitempty"`
	Pool *Pool           `json:"pool,omitempty"`
}

// Message - struct to keep messages
type Message struct {
	Type  int    `json:"type,omitempty"`
	Body  string `json:"body,omitempty"`
	UsrID int    `json:"usr_id,omitempty"`
}

//Read - method to read client messages
func (c *Client) Read() {
	defer func() {
		log.Printf("connetion is closed %v", c)
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//best is to send json - simplify using map - but also more flexible
		pMap := make(map[string]interface{})

		err = json.Unmarshal(p, &pMap)
		if err != nil {
			log.Println(err)
			// return
		}
		log.Printf("body is %+v", pMap)

		data := pMap["value"].(string)
		ID := int(pMap["ID"].(float64)) //default is float 64 as number form front-end, switch statement could be used

		message := Message{Type: messageType, Body: data, UsrID: ID}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
