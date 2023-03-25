package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Client struct {
	conn     *websocket.Conn
	nickname string
}

type Message struct {
	Nickname string `json:"nickname"`
	Text     string `json:"text"`
}

var clients = make(map[*websocket.Conn]*Client) // Conn is key, *Client is value

func (c *Client) readFromClient() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			delete(clients, c.conn)
			return
		}
		messageWithNickname := Message{Nickname: c.nickname, Text: string(message)}
		broadcast(messageWithNickname)
	}

}

func broadcast(message Message) {
	marshaledMessage, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	for conn, _ := range clients {
		err := conn.WriteMessage(websocket.TextMessage, marshaledMessage)
		if err != nil {
			delete(clients, conn)
			continue
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:3000"
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketHandler(c *gin.Context) {
	fmt.Println("HERE")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//defer func() {
	//	conn.Close()
	//	delete(clients, conn)
	//}()

	var nickname string
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			return
		}
		if msg["type"] == "userJoin" {
			nickname = msg["nickname"].(string)
			break
		}
	}

	client := &Client{conn: conn, nickname: nickname}
	clients[conn] = client

	go client.readFromClient()

}

func main() {
	r := gin.Default()
	r.GET("/", socketHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
