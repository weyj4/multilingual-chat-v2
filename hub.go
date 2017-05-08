package main

import "fmt"

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
	languages  map[string]bool
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		languages:  make(map[string]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			if _, ok := h.languages[client.language]; !ok {
				h.languages[client.language] = true
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			translations := make(map[string][]byte)
			for language := range h.languages {
				if message.language != language {
					fmt.Println("message.language", message.language)
					fmt.Println("language", language)
					translations[language] = Translate(message.language, language, string(message.msg))
					for k, v := range translations {
						fmt.Println("key", k, "value", string(v))
					}
				} else {
					translations[language] = message.msg
				}
			}
			for client := range h.clients {
				select {
				case client.send <- translations[client.language]:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
