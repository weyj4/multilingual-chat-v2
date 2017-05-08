package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	nick     []byte
	language string
}

type Message struct {
	msg      []byte
	language string
}

func (client *Client) readPump() {
	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		client.hub.broadcast <- Message{
			msg:      message,
			language: client.language,
		}
	}
}

func (client *Client) writePump() {
	for {
		message, ok := <-client.send
		if !ok {
			return
		}

		w, err := client.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)

		n := len(client.send)
		for i := 0; i < n; i++ {
			w.Write(newline)
			w.Write(<-client.send)
		}

		if err := w.Close(); err != nil {
			return
		}
	}
}

func promptNick(client *Client) {
	fmt.Println("promptNick")
	client.conn.WriteMessage(websocket.TextMessage, []byte("Welcome to the multilingual chat! What is your nickname?"))

	fmt.Println("reading")
	_, nick, err := client.conn.ReadMessage()
	if err != nil {
		return
	}
	client.nick = append(nick, []byte(": ")...)

	client.conn.WriteMessage(websocket.TextMessage, append([]byte("Hello, "), nick...))
}

func promptLang(client *Client) {
	client.conn.WriteMessage(websocket.TextMessage, []byte("Pick a language"))

	_, lang, err := client.conn.ReadMessage()
	if err != nil {
		return
	}
	client.language = string(lang)
}

func HandleClient(w http.ResponseWriter, r *http.Request, hub *Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte),
	}
	_, init, err := client.conn.ReadMessage()
	if err != nil {
		return
	}
	fmt.Println("init", init)
	promptNick(client)
	promptLang(client)

	fmt.Println(client)

	client.hub.register <- client
	go client.writePump()
	client.readPump()
}
